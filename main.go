// main.go
package main

import (
	"flocknest/app/middleware"
	"flocknest/app/routes"
	"flocknest/app/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := utils.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	// API-2
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// API-1
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
