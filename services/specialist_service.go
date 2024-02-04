package services

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Dubbril/go-edge-utility/models"
	"os"
	"runtime"
	"strconv"
	"strings"
)

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
