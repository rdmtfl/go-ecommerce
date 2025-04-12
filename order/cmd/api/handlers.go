package main

import (
	"net/http"

	"github.com/rdmtfl/order/data"
)

// WriteLog -
func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload data.OrderEntry
	_ = app.readJson(w, r, &requestPayload)

	// send to db
	err := app.Models.OrderEntry.Insert(requestPayload)
	if err != nil {
		app.errorJson(w, err)
		return
	}

	// report user
	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJson(w, http.StatusAccepted, resp)
}
