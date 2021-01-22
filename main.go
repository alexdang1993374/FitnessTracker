package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/routes"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var x int = 0

func main() {
	config.Connect()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./fitness-tracker/build", true)))

	// Setup route group for the API
	routes.Routes(router)

	// Start and run the server
	router.Run(":5000")
}
