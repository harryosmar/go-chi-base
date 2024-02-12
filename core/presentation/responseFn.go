package presentation

import (
	"github.com/go-playground/validator/v10"
	"github.com/harryosmar/go-chi-base/core/errors"
	"net/http"
)

func ResponseErr(w http.ResponseWriter, err error) {
	var respErr codes.CodeErrEntity
	if codeErr, ok := err.(codes.CodeErrEntity); ok {
		respErr = codeErr
	} else {
		respErr = codes.NewCodeErrf(codes.ErrGeneral).SetDetail(err.Error())
	}

	NewResponseEntity().
		WithStatusCode(respErr.HttpStatus()).
		WithContentStatus(false).
		WithMessage(respErr.String()).
		WithErrorCode(respErr.Code()).
		WithErrorDetail(respErr.Detail()).
		WriteJson(w)
}

func ResponseOk(w http.ResponseWriter, statusCode int, data interface{}) {
	NewResponseEntity().
		WithStatusCode(statusCode).
		WithContentStatus(true).
		WithData(data).
		WriteJson(w)
}

func ResponseErrValidation(w http.ResponseWriter, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		data := map[string]interface{}{}
		for _, err := range errors {
			data[err.Field()] = err.Error()
		}

		NewResponseEntity().
			WithStatusCode(codes.ErrValidation.HttpStatus()).
			WithContentStatus(false).
			WithMessage(codes.ErrValidation.String()).
			WithErrorCode(codes.ErrValidation.Code()).
			WithData(data).
			WriteJson(w)

		return
	}

	ResponseErr(w, err)
}
