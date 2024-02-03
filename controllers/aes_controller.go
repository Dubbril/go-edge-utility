package controllers

import (
	"github.com/Dubbril/go-edge-utility/models"
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AesController struct {
	AesService *services.AesService
}

func NewAesController(aesService *services.AesService) *AesController {
	return &AesController{AesService: aesService}
}

func (ctrl *AesController) Encrypt(c *gin.Context) {

	// Bine @RequestBody @Valid AesRequest aesRequest in java
	var aesReq models.AesRequest
	if err := c.ShouldBindJSON(&aesReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.AesService.EncryptData(&aesReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, data)
}

func (ctrl *AesController) Decrypt(c *gin.Context) {
	var aesReq models.AesRequest
	if err := c.ShouldBindJSON(&aesReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ctrl.AesService.DecryptData(&aesReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, data)
}
