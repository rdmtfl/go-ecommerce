package main

import (
	"net/http"
	"time"

	"github.com/rdmtfl/order/data"
)

type RequestPayload struct {
	OrderID      string    `json:"orderId"`
	CustomerID   string    `json:"customerId"`
	OrderDate    time.Time `json:"orderDate"`
	Status       string    `json:"status"`
	Subtotal     float64   `json:"subtotal"`
	Tax          float64   `json:"tax"`
	ShippingCost float64   `json:"shippingCost"`
	Total        float64   `json:"total"`
}

// WriteLog -
func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload RequestPayload
	_ = app.readJson(w, r, &requestPayload)

	// insert data
	event := data.OrderEntry{
		OrderID:    requestPayload.OrderID,
		CustomerID: requestPayload.CustomerID,
		OrderDate:  requestPayload.OrderDate,
		Status:     requestPayload.Status,
		//Items:           requestPayload.Items,
		Subtotal:     requestPayload.Subtotal,
		Tax:          requestPayload.Tax,
		ShippingCost: requestPayload.ShippingCost,
		Total:        requestPayload.Total,
		//ShippingMethod: requestPayload.ShippingMethod,
		//PaymentMethod:  requestPayload.PaymentMethod,
		//BillingAddress:  requestPayload.BillingAddress,
		//ShippingAddress: requestPayload.ShippingAddress,
		//Notes: requestPayload.Notes,
		//Discounts:       requestPayload.Discounts,
		//CreatedAt: requestPayload.CreatedAt,
		//UpdatedAt: requestPayload.UpdatedAt,
	}

	err := app.Models.OrderEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJson(w, http.StatusAccepted, resp)
}
