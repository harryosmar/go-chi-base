package validator

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func ValidateRequest[K any](r *http.Request, s *K) error {
	if validate == nil {
		validate = validator.New()
	}

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		return err
	}

	return validate.Struct(s)
}
