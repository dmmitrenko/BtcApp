package main

import (
	"example/BtcApp/handlers"
	"log"

	_ "example/BtcApp/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GSES2 BTC application
// @version 1.0
// @description API documentation for GSES2 BTC application
// @BasePath /
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/exchange-rate", handlers.GetExchangeRate)
	router.POST("/email", handlers.SubscribeEmail)
	router.POST("/subscribe", handlers.SendEmails)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
