package main

import (
	"fmt"
	"github.com/Dubbril/go-edge-utility/config"
	"github.com/Dubbril/go-edge-utility/controllers"
	"github.com/Dubbril/go-edge-utility/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {

	// Initialize viper
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// Set Gin to release mode to disable debug output
	gin.SetMode(gin.ReleaseMode)
	// Create a new Gin router
	router := gin.Default()
	config.InitHomePage(router)

	// Handle Log Request & Response
	//router.Use(middleware.LogHandler())

	// Register Aes Controller
	aesService := services.NewAesService()
	aesController := controllers.NewAesController(aesService)
	router.POST("/api/v1/cryptography/aes/encrypt", aesController.Encrypt)
	router.POST("/api/v1/cryptography/aes/decrypt", aesController.Decrypt)

	// Register Json Controller
	jsonService := services.NewJsonService()
	jsonController := controllers.NewJsonController(jsonService)
	router.POST("/api/v1/cryptography/json/escaped", jsonController.EscapedJSON)
	router.POST("/api/v1/cryptography/json/base64/image", jsonController.Base64ToImage)
	router.POST("/api/v1/cryptography/json/image/base64", jsonController.ImageToBase64)
	router.POST("/api/v1/cryptography/json/snake-to-camel", jsonController.JSONSnakeToCamel)
	router.POST("/api/v1/cryptography/json/camel-to-snake", jsonController.JSONCamelToSnake)

	// Register Specialist Controller
	specialistService := services.NewSpecialistService()
	specialistController := controllers.NewSpecialistController(specialistService)
	router.GET("/api/v1/cryptography/specialist/download", specialistController.DownloadFile)
	router.POST("/api/v1/cryptography/specialist/make", specialistController.MakeSpecialist)
	router.POST("/api/v1/cryptography/specialist/delete", specialistController.DeleteSpecialist)
	router.GET("/api/v1/cryptography/specialist/query", specialistController.FindByCustomerNo)
	router.GET("/api/v1/cryptography/specialist/delete-by-index", specialistController.DeleteByIndex)

	// Open browser on start
	config.OpenBrowser("http://localhost:8083")

	// Run the server on port 8080
	err := router.Run(":8083")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start on port 8083")
		return
	}
}
