package controllers

import (
	"bytes"
	"github.com/Dubbril/go-edge-utility/services"
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
