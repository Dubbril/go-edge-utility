package main

import (
	"github.com/Dubbril/go-edge-utility/controllers"
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")
	router.Static("/css", "static/css")
	router.Static("/js", "static/js")

	// Set up routes
	homeController := controllers.NewHomeController()
	router.GET("/", homeController.Index)
	router.GET("/favicon.ico", homeController.FaviconHandler)

	aesService := services.NewAesService()
	aesController := controllers.NewAesController(aesService)
	router.POST("/api/v1/cryptography/aes/encrypt", aesController.Encrypt)
	router.POST("/api/v1/cryptography/aes/decrypt", aesController.Decrypt)

	jsonService := services.NewJsonService()
	jsonController := controllers.NewJsonController(jsonService)
	router.POST("/api/v1/cryptography/json/escaped", jsonController.EscapedJSON)

	// Run the server on port 8080
	err := router.Run(":8083")
	if err != nil {
		return
	}
}
