package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeResponse(w, struct {
		Message  any
		Code     int
		Error    bool
		Body     interface{}
		MetaData interface{}
	}{
		Message:  "Health Status",
		Code:     200,
		Error:    false,
		Body:     data,
		MetaData: nil,
	})
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
