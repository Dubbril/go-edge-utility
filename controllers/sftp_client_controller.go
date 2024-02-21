package controllers

import (
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SftpClientController struct {
	SftpClientService *services.SftpClientService
}

func NewSftpClientController(specialistService *services.SftpClientService) *SftpClientController {
	return &SftpClientController{SftpClientService: specialistService}
}

func (ctrl *SftpClientController) DownloadLastFileOfSpecialist(c *gin.Context) {
	result, err := ctrl.SftpClientService.DownloadLastFileOfSpecialist()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"specialistFileName": result})
}
