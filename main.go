package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	SetupEnv()

	router := setupRouter()

	port := getConfig("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

// setupRouter starts gin web server with setting router.
func setupRouter() *gin.Engine {
	mode := getConfig("GIN_MODE")

	gin.SetMode(mode)

	// Initialize gin router
	router := gin.New()

	// Define routes
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	router.GET("/pantry", getItems)
	router.GET("/pantry/:id", getItemById)
	router.POST("/item", postItems)

	return router
}
