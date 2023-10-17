package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type item struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Amount         string    `json:"amount"`
	IsLow          bool      `json:"isLow"`
	ExpirationDate time.Time `json:"expirationDate"`
	PurchaseDate   time.Time `json:"purchaseDate"`
}

var pantry = []item{
	{ID: "1", Name: "Milk", Amount: "1 gallon", IsLow: false, ExpirationDate: time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), PurchaseDate: time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC)},
}

func main() {
	router := gin.Default() // initialize gin router

	router.GET("/pantry", getItems)
	router.GET("/pantry/:id", getItemById)
	router.POST("/item", postItems)

	router.Run("localhost:8080") // starts a http server
}

// getItems responds with a list of all pantry items as JSON
func getItems(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, pantry)
}

// postItems creates a new item and responds with the added item as JSON
func postItems(context *gin.Context) {
	var newItem item

	err := context.BindJSON(&newItem) // binds the received json to the newItem

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "err": err})
	} else {
		pantry = append(pantry, newItem)

		context.IndentedJSON(http.StatusCreated, newItem)
	}
}

// getItemById responds with the item corresponding to the ID url parameter as JSON
func getItemById(context *gin.Context) {
	id := context.Param("id")

	for _, i := range pantry {
		if i.ID == id {
			context.IndentedJSON(http.StatusAccepted, i)
			return
		}
	}

	context.IndentedJSON(http.StatusNoContent, gin.H{"message": "Item not found"})
}
