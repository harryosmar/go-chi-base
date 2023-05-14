package presentation

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/harryosmar/go-chi-base/errors"
	"net/http"
)

type ResponseEntity struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	ErrorCode   string      `json:"error_code"`
	ErrorDetail string      `json:"error_detail"`
	Data        interface{} `json:"data"`
}

func ResponseErr(w http.ResponseWriter, err error) {
	var respErr codes.CodeErrEntity
	if codeErr, ok := err.(codes.CodeErrEntity); ok {
		respErr = codeErr
	} else {
		respErr = codes.NewCodeErrf(codes.ErrGeneral).SetDetail(err.Error())
	}

	w.WriteHeader(respErr.HttpStatus())
	json.NewEncoder(w).Encode(ResponseEntity{
		Message:     respErr.String(),
		ErrorCode:   respErr.Code(),
		ErrorDetail: respErr.Detail(),
	})
}

func Response(w http.ResponseWriter, statusCode int, data interface{}) {
	statusText := http.StatusText(statusCode)
	if statusText == "" {
		ResponseErr(w, codes.NewCodeErrf(codes.ErrInvalidStatusCode).Paramf(statusCode))
		return
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseEntity{
		Status: true,
		Data:   data,
	})
}

func ResponseErrValidation(w http.ResponseWriter, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		w.WriteHeader(codes.ErrValidation.HttpStatus())
		data := map[string]interface{}{}
		for _, err := range errors {
			data[err.Field()] = err.Error()
		}
		json.NewEncoder(w).Encode(ResponseEntity{
			Message:   codes.ErrValidation.String(),
			ErrorCode: codes.ErrValidation.Code(),
			Data:      data,
		})
		return
	}

	ResponseErr(w, err)
}
