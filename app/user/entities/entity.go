package entities

import "time"

type ValidateCredentialRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ValidateCredentialResponse struct {
	Message string  `json:"message"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}
