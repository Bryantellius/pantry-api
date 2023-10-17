package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupEnv()

	SetupServer()
}

// SetupServer starts gin web server with setting router.
func SetupServer() {
	mode := getConfig("GIN_MODE")

	gin.SetMode(mode)

	router := gin.New()

	router.GET("/pantry", getItems)
	router.GET("/pantry/:id", getItemById)
	router.POST("/item", postItems)

	port := getConfig("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
