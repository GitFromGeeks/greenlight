package main

import (
	"errors"
	"net/http"

	"greenlight.altamash.dev/internal/data"
	"greenlight.altamash.dev/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Users.Insert(user)
	if err != nil {
		switch {

		case errors.Is(err, data.ErrDuplicateEmail):
			app.logger.Error("Duplicated Error")
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.logger.Error("Default Error")

			app.serverErrorResponse(w, r, err)
		}
		return

	}

	err = app.writeResponse(w, struct {
		Message  any
		Code     int
		Error    bool
		Body     interface{}
		MetaData interface{}
	}{
		Message:  "User Created Successfully",
		Code:     http.StatusOK,
		Error:    false,
		Body:     user,
		MetaData: nil,
	})
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
