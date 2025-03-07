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
		Message string
		Code    int
		Error   bool
		Body    interface{}
	}{
		Message: "Health Status",
		Code:    200,
		Error:   false,
		Body:    data,
	})
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
