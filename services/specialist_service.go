package services

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Dubbril/go-edge-utility/models"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var dataSpecialist *[]models.SpecialistRequest

type SpecialistService struct{}

func NewSpecialistService() *SpecialistService {
	return &SpecialistService{}
}

func (s SpecialistService) getDataFromExcel(fileBytes []byte) (map[int][]string, error) {
	xlsx, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return nil, fmt.Errorf("error opening Excel file: %w", err)
	}

	data := make(map[int][]string)

	// Get sheet name from the first sheet in the workbook
	sheetName := xlsx.GetSheetName(1)

	rows := xlsx.GetRows(sheetName)
	rows = rows[1:]
	for i, row := range rows {
		row = row[1:]
		txtList := make([]string, len(row))
		for j := range row {
			cell := xlsx.GetCellValue(sheetName, excelize.ToAlphaString(j+1)+strconv.Itoa(i+2))
			txtList[j] = cell
		}
		data[i+1] = txtList
	}

	return data, nil
}

func (s SpecialistService) MakeSpecialist(fileBytes []byte, index int) (string, error) {
	dataFromExcel, err := s.getDataFromExcel(fileBytes)
	if err != nil {
		return "", fmt.Errorf("error getting data from Excel: %w", err)
	}

	currentIndex := index
	var result strings.Builder
	for _, value := range dataFromExcel {
		customerId := value[0]
		customerName := value[1]
		specialCode := value[2][len(value[2])-2:]

		strResult := fmt.Sprintf("%d|%s|%s|%s||||TH|ZZ|0|R|B|Y||%s|%s|(1)(2)(3)"+SystemLineSeparator(),
			currentIndex, customerId, customerName, customerName, specialCode, customerId)

		result.WriteString(strResult)
		currentIndex++
	}

	return strings.TrimSpace(result.String()), nil
}

// SystemLineSeparator provides a platform-independent line separator
func SystemLineSeparator() string {
	switch runtime.GOOS {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}

func (s SpecialistService) DownloadExampleFile() ([]byte, error) {
	// Open the file from the example package
	file, err := os.ReadFile("example/Add_Special_list_Datamart.xlsx")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return file, nil
}

func (s SpecialistService) DeleteSpecialist(fileContent []byte, deleteIDs map[string]struct{}) ([]byte, error) {
	reader := bytes.NewReader(fileContent)
	scanner := bufio.NewScanner(reader)
	var index int
	var result strings.Builder

	// Read the header
	if scanner.Scan() {
		header := scanner.Text()
		result.WriteString(header)
	}

	// Process the data
	for scanner.Scan() {
		dataLine := scanner.Text()
		specialist := models.NewSpecialistRequest(dataLine)

		if _, shouldDelete := deleteIDs[specialist.CustomerNo]; shouldDelete {
			continue
		}
		specialist.RowNo = fmt.Sprint(index + 1)
		result.WriteString(SystemLineSeparator())
		result.WriteString(specialist.String())
		index++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file content: %w", err)
	}

	return []byte(result.String()), nil
}

func paginate(specialistSlice []models.SpecialistRequest, page int) ([]models.SpecialistRequest, int) {
	pageSize := 20
	totalRecords := len(specialistSlice)
	totalPages := (totalRecords + pageSize - 1) / pageSize

	startIndex := (page - 1) * pageSize
	endIndex := page * pageSize
	if endIndex > totalRecords {
		endIndex = totalRecords
	}

	return specialistSlice[startIndex:endIndex], totalPages
}

func (s SpecialistService) DeleteByIndex(index int) error {
	if !isEmpty(dataSpecialist) {
		// Check if the index is valid
		if index < 0 || index >= len(*dataSpecialist) {
			return errors.New("index out of range")
		}

		var slice = *dataSpecialist

		// Use append to create a new slice excluding the element at the specified index
		data := append(slice[:index], slice[index+1:]...)
		dataSpecialist = &data
	}

	return nil
}

func (s SpecialistService) FilterByCustomerNo(customerNo string) ([]interface{}, error) {
	err := readSpecialistFromFile(sftpInfo.Path)
	if err != nil {
		return nil, err
	}

	var responseData []interface{}
	for index, value := range *dataSpecialist {
		if customerNo == value.CustomerNo {
			data := map[string]interface{}{
				"index": index,
				"data":  value,
			}
			responseData = append(responseData, data)

		}
	}

	return responseData, nil
}

func isEmpty(s *[]models.SpecialistRequest) bool {
	return s == nil || len(*s) == 0
}

func readSpecialistFromFile(filePath string) error {

	if !isEmpty(dataSpecialist) {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read and discard the first line

	var specialistReqSlice []models.SpecialistRequest
	for scanner.Scan() {
		readLine := scanner.Text()
		specialistReq := models.NewSpecialistRequest(readLine)
		specialistReqSlice = append(specialistReqSlice, *specialistReq)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	dataSpecialist = &specialistReqSlice

	return nil
}
