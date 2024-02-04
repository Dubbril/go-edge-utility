package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/Dubbril/go-edge-utility/utils"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type JsonController struct {
	JsonService *services.JsonService
}

func NewJsonController(jsonService *services.JsonService) *JsonController {
	return &JsonController{JsonService: jsonService}
}

func (ctrl *JsonController) EscapedJSON(c *gin.Context) {

	// Bine @RequestBody String strJson in java
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Convert the buffer to a string
	requestBody := buffer.String()
	result, err := ctrl.JsonService.EscapedJSON(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, result)
}

func (ctrl *JsonController) Base64ToImage(c *gin.Context) {

	// Bine @RequestBody String strJson in java
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Convert the buffer to a string
	requestBody := buffer.String()
	result, err := ctrl.JsonService.Base64ToImage(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contentType := utils.CheckImageType(result)
	c.Data(http.StatusOK, string(contentType), result)
}

func (ctrl *JsonController) ImageToBase64(c *gin.Context) {

	// Bine @RequestBody String strJson in java
	var buffer bytes.Buffer
	_, err := io.Copy(&buffer, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	// Convert the buffer to byte[]
	requestBody := buffer.Bytes()
	result := ctrl.JsonService.ImageToBase64(requestBody)

	c.String(http.StatusOK, result)
}

func (ctrl *JsonController) JSONSnakeToCamel(c *gin.Context) {

	var requestBody map[string]interface{}

	// Bind request body to a map
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	camelCaseObject := ctrl.JsonService.JSONSnakeToCamel(requestBody)

	// Convert back to JSON for display
	camelCaseJSON, err := json.MarshalIndent(camelCaseObject, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	c.Data(http.StatusOK, "application/json", camelCaseJSON)
}

func (ctrl *JsonController) JSONCamelToSnake(c *gin.Context) {

	var requestBody map[string]interface{}

	// Bind request body to a map
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	snakeCaseObject := ctrl.JsonService.JSONCamelToSnake(requestBody)

	// Convert back to JSON for display
	snakeCaseJSON, err := json.MarshalIndent(snakeCaseObject, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	c.Data(http.StatusOK, "application/json", snakeCaseJSON)
}
