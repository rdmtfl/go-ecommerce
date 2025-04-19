package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rdmtfl/order/data"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJson
func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // 1 mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// writeJson
func (app *Config) writeJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

// errorJson
func (app *Config) errorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJson(w, statusCode, payload)
}

// validatePayload
func (app *Config) validatePayload(payload data.OrderEntry) (args []string) {
	args = make([]string, 0)

	// order
	if strings.EqualFold(payload.OrderID, "") {
		args = append(args, "Order ID is required")
	}
	if strings.EqualFold(payload.CustomerID, "") {
		args = append(args, "Customer ID is required")
	}
	if strings.EqualFold(payload.Status, "") {
		args = append(args, "Status is required")
	}
	if payload.Subtotal <= 0 {
		args = append(args, "Subtotal must be greater than 0")
	}
	if payload.Total <= 0 {
		args = append(args, "Total must be greater than 0")
	} else if payload.Total != payload.Subtotal+payload.Tax+payload.ShippingCost {
		args = append(args, "Total must be equal to subtotal + tax + shipping cost")
	}
	if strings.EqualFold(payload.ShippingMethod, "") {
		args = append(args, "Shipping method is required")
	}
	if strings.EqualFold(payload.PaymentMethod, "") {
		args = append(args, "Payment method is required")
	}

	// // order items
	if len(payload.Items) == 0 {
		args = append(args, "At least one item is required")
	} else {
		for index, item := range payload.Items {
			if strings.EqualFold(item.ProductID, "") {
				args = append(args, fmt.Sprintf("Product ID is required for item %d", index))
			}
			if item.Quantity <= 0 {
				args = append(args, fmt.Sprintf("Quantity must be greater than 0 for item %d", index))
			}
			if item.UnitPrice <= 0 {
				args = append(args, fmt.Sprintf("Unit Price must be greater than 0 for item %d", index))

			}
			if strings.EqualFold(item.SKU, "") {
				args = append(args, fmt.Sprintf("SKU is required for item %d", index))
			}
		}
	}

	// billing address
	if strings.EqualFold(payload.BillingAddress.PostalCode, "") {
		args = append(args, "Billing address postal code is required")
	}
	// shipping address
	if strings.EqualFold(payload.ShippingAddress.PostalCode, "") {
		args = append(args, "Shipping address postal code is required")
	}

	// order discounts
	var sumOfDiscounts float64
	if payload.Discounts != nil {
		for index, discount := range payload.Discounts {
			sumOfDiscounts += discount.Amount
			if strings.EqualFold(discount.Code, "") {
				args = append(args, fmt.Sprintf("Code required for discount %d", index))
			}

			if discount.Amount <= 0 {
				args = append(args, fmt.Sprintf("Amount must be greater than 0 for discount %d", index))
			}

			if discount.IsPercent {
				if discount.Amount > 100 {
					args = append(args, fmt.Sprintf("Percentage cannot be greater than 100 for discount %d", index))
				}
			}
		}

		if sumOfDiscounts > payload.Subtotal {
			args = append(args, "Sum of discounts cannot be greater than subtotal")
		}
	}

	return args
}
