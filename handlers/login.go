package handlers

import (
	"github.com/harryosmar/go-chi-base/presentation"
	validator2 "github.com/harryosmar/go-chi-base/validator"
	"net/http"
)

type Credentials struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	defer r.Body.Close()

	if err := validator2.ValidateRequest(r, credentials); err != nil {
		presentation.ResponseErrValidation(w, err)
		return
	}

	// Perform authentication logic here
	// ...

	// Example response
	response := map[string]string{"message": "Login successful"}
	presentation.Response(w, 200, response)

	//presentation.ResponseErr(w, errors.New("abc"))

}
