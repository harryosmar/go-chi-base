package handlers

import (
	"github.com/harryosmar/go-chi-base/app/user/entities"
	"github.com/harryosmar/go-chi-base/presentation"
	validator2 "github.com/harryosmar/go-chi-base/validator"
	"net/http"
)

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var credentials entities.ValidateCredentialRequest
	if err := validator2.ValidateRequest(r, &credentials); err != nil {
		presentation.ResponseErrValidation(w, err)
		return
	}

	validateResult, err := u.service.ValidateCredentials(r.Context(), credentials)
	if err != nil {
		presentation.ResponseErr(w, err)
		return
	}

	presentation.ResponseOk(w, 200, validateResult)
}
