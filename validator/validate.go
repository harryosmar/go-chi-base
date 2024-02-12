package validator

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	codes "github.com/harryosmar/go-chi-base/errors"
	"net/http"
)

var validate *validator.Validate

func ValidateRequest[K any](r *http.Request, s *K) error {
	if validate == nil {
		validate = validator.New()
	}

	body := r.Body
	if r.Body == http.NoBody {
		return codes.ErrValidationEmptyRequestBody
	}

	if err := json.NewDecoder(body).Decode(&s); err != nil {
		return err
	}

	return validate.Struct(s)
}
