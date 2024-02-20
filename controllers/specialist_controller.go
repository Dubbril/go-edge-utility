package controllers

import (
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/Dubbril/go-edge-utility/utils"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type SpecialistController struct {
	SpecialistService *services.SpecialistService
}

func NewSpecialistController(specialistService *services.SpecialistService) *SpecialistController {
	return &SpecialistController{SpecialistService: specialistService}
}

func (ctrl *SpecialistController) MakeSpecialist(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing file parameter"})
		return
	}

	// Open the file
	fileHandle, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {
			return
		}
	}(fileHandle)

	// Read the file content into a byte slice
	fileContent, err := io.ReadAll(fileHandle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	// Retrieve the "index" parameter from form data
	indexStr := c.PostForm("index")
	if indexStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing index parameter"})
		return
	}

	// Convert index to an integer
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index parameter"})
		return
	}

	result, err := ctrl.SpecialistService.MakeSpecialist(fileContent, index)
	if err != nil {
		return
	}

	c.String(http.StatusOK, result)
}

func (ctrl *SpecialistController) DeleteSpecialist(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing file parameter"})
		return
	}

	// Open the file
	fileHandle, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer func(fileHandle multipart.File) {
		err := fileHandle.Close()
		if err != nil {
			return
		}
	}(fileHandle)

	// Read the file content into a byte slice
	fileContent, err := io.ReadAll(fileHandle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	deleteId := c.PostForm("deleteId")
	// Split the deleteId string into individual IDs
	idList := strings.Split(deleteId, ",")
	// Create a set-like structure using a map[string]struct{}
	deleteIDSet := make(map[string]struct{})
	// Add each ID to the set
	for _, id := range idList {
		deleteIDSet[id] = struct{}{}
	}

	if deleteId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing index parameter"})
		return
	}

	result, err := ctrl.SpecialistService.DeleteSpecialist(fileContent, deleteIDSet)
	if err != nil {
		return
	}

	fileName := "EIM_EDGE_BLACKLIST_" + utils.DateWithYearMonthDay() + ".txt"

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/octet-stream", result)
}

func (ctrl *SpecialistController) DownloadFile(c *gin.Context) {
	// Download the file content
	fileContent, err := ctrl.SpecialistService.DownloadExampleFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download file"})
		return
	}

	// Set response headers
	c.Header("Content-Disposition", "attachment; filename=Add_Special_list_Datamart.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", fileContent)
}

func (ctrl *SpecialistController) FindByCustomerNo(c *gin.Context) {
	customerNoFilter := c.Query("customerNo")
	if customerNoFilter == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CustomerNo is empty"})
	}

	specialistData, err := ctrl.SpecialistService.FilterByCustomerNo(customerNoFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, specialistData)
}

func (ctrl *SpecialistController) DeleteByIndex(c *gin.Context) {
	index, err := strconv.Atoi(c.Query("index"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	err = ctrl.SpecialistService.DeleteByIndex(index)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Delete Success"})
}
