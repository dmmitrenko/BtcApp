package main

import (
	_ "CurrencyRateApp/api/docs"
	"CurrencyRateApp/api/route"
)

// @title GSES2 BTC application
// @version 1.0
// @description API documentation for GSES2 BTC application
// @BasePath /
func main() {
	router := route.SetupRouter()
	router.Run(":8080")
}
