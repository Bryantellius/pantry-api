package main

import "time"

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
