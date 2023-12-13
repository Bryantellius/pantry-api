package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var router = setupRouter()

func TestGetItems(t *testing.T) {
	w := httptest.NewRecorder()
	
	req, _ := http.NewRequest("GET", "/pantry", nil)
	router.ServeHTTP(w, req)

	// Checks for correct status code
	assert.Equal(t, http.StatusOK, w.Code, "Should return ok status code")

	var jsonResponse *[]item

	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

	if err != nil {
		panic("Cannot parse json response")
	}

	// Checks for correct response body
	assert.Equal(t, &pantry, jsonResponse, "Should return list of pantry items")
}

func TestGetItem(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pantry/1", nil)
	router.ServeHTTP(w, req)

	// Check for the correct status code
	assert.Equal(t, http.StatusAccepted, w.Code, "Should return accepted status code")

	var jsonResponse *item

	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)

	if err != nil {
		panic("Cannot parse json response")
	}

	// Check for correct response body
	assert.Equal(t, &pantry[0], jsonResponse, "Should return correct pantry item")
}

func TestPostItem(t *testing.T) {
	w := httptest.NewRecorder()

	newItem := item{ID: "2", Name: "Test Name", Amount: "Test amount", IsLow: false, ExpirationDate: time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), PurchaseDate: time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC)}
	b, _ := json.Marshal(newItem)

	req, _ := http.NewRequest("POST", "/item", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	// Check for the correct status code
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var actualItem *item 

	json.Unmarshal(w.Body.Bytes(), &actualItem)

	// Check for correct response body
	assert.Equal(t, &newItem, actualItem)
}
