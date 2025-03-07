package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeResponse(w http.ResponseWriter, params struct {
	Message string
	Code    int
	Error   bool
	Body    interface{}
}) error {
	response := Response{}
	response.Headers.Message = params.Message
	response.Headers.Code = params.Code
	response.Headers.Error = params.Error
	response.Body = params.Body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(params.Code)
	json.NewEncoder(w).Encode(response)
	return nil
}

type Response struct {
	Headers struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Error   bool   `json:"error"`
	} `json:"headers"`
	Body interface{} `json:"body"`
}
