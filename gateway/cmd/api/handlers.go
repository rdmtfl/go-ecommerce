package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/rdmtfl/gateway/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func (app *Config) Gateway(w http.ResponseWriter, r *http.Request) {
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
		app.postOrder(w, requestPayload.Order)
	default:
		app.errorJson(w, errors.New("unknown action"))
	}
}

// postOrder -
func (app *Config) postOrder(w http.ResponseWriter, entry OrderPayload) {
	jsonData, err := json.Marshal(entry)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	logServiceURL := "http://order-service/order"

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

	// read response from service
	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	payload := jsonResponse{
		Error:   jsonFromService.Error,
		Message: jsonFromService.Message,
		Data:    jsonFromService.Data,
	}
	app.writeJson(w, response.StatusCode, payload)
}

// PostOrderViaGRPC -
func (app *Config) PostOrderViaGRPC(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	conn, err := grpc.Dial("order-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.errorJson(w, err)
		return
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.PostOrder(ctx, &pb.OrderRequest{
		OrderEntry: &pb.Order{
			OrderId:    requestPayload.Order.OrderID,
			CustomerId: requestPayload.Order.CustomerID,
		},
	})

	if err != nil {
		app.errorJson(w, err)
		return
	}

	app.writeJson(w, http.StatusAccepted, res)

	// var payload jsonResponse
	// payload.Error = false
	// payload.Message = "logged"

	// app.writeJson(w, http.StatusAccepted, payload)
}
