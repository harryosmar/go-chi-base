package entities

type ValidateCredentialRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ValidateCredentialResponse struct {
	Message string  `json:"message"`
	Profile Profile `json:"profile"`
}
