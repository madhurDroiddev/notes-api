package main

import (
	"notes-api/config"

	"notes-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Notes API is alive"})
	})

	r.Run(":8080")
}
