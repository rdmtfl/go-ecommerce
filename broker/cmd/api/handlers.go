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
	OrderID      string    `json:"orderId"`
	CustomerID   string    `json:"customerId"`
	OrderDate    time.Time `json:"orderDate"`
	Status       string    `json:"status"`
	Subtotal     float64   `json:"subtotal"`
	Tax          float64   `json:"tax"`
	ShippingCost float64   `json:"shippingCost"`
	Total        float64   `json:"total"`
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
