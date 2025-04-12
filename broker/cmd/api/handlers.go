package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type RequestPayload struct {
	Action string       `json:"action"`
	Order  OrderPayload `json:"order,omitempty"`
}

type OrderPayload struct {
	OrderID         string      `json:"orderId"`
	CustomerID      string      `json:"customerId"`
	OrderDate       time.Time   `json:"orderDate"`
	Status          string      `json:"status"`
	Items           []OrderItem `json:"items"`
	Subtotal        float64     `json:"subtotal"`
	Tax             float64     `json:"tax"`
	ShippingCost    float64     `json:"shippingCost"`
	Total           float64     `json:"total"`
	ShippingMethod  string      `json:"shippingMethod"`
	PaymentMethod   string      `json:"paymentMethod"`
	BillingAddress  Address     `json:"billingAddress"`
	ShippingAddress Address     `json:"shippingAddress"`
	Notes           string      `json:"notes,omitempty"`
	Discounts       []Discount  `json:"discounts,omitempty"`
}

type OrderItem struct {
	ProductID   string  `json:"productId"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Discount struct {
	Code        string  `json:"code"`
	Amount      float64 `json:"amount"`
	IsPercent   bool    `json:"isPercent"`
	Description string  `json:"description"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "connected!",
	}

	_ = app.writeJson(w, http.StatusOK, payload)
}

// HandleSubmission
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	switch requestPayload.Action {
	case "order":
		app.logItem(w, requestPayload.Order)
	default:
		app.errorJson(w, errors.New("unknown action"))
	}
}

func (app *Config) logItem(w http.ResponseWriter, entry OrderPayload) {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://order-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJson(w, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJson(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJson(w, http.StatusAccepted, payload)

}
