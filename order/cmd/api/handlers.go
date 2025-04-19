package main

import (
	"net/http"
	"time"

	"github.com/rdmtfl/order/data"
)

type responsePayload struct {
	OrderId      string    `json:"orderId"`
	CreationDate time.Time `json:"creationDate"`
}

// WriteLog -
func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.OrderEntry
	var payload jsonResponse

	// decode request
	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// validate payload
	args := app.validatePayload(requestPayload)
	if len(args) > 0 {
		payload.Error = true
		payload.Message = "missing fields"
		payload.Data = args

		app.writeJson(w, http.StatusBadRequest, payload)
		return
	}

	// send to db
	err = app.Models.OrderEntry.Insert(requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// done correctly! report success
	payload.Error = false
	payload.Message = "created"
	payload.Data = responsePayload{
		OrderId:      requestPayload.OrderID,
		CreationDate: time.Now(),
	}

	app.writeJson(w, http.StatusAccepted, payload)
}
