package entities

type AuthenticateRequest struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

type ValidateTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

type JwtClaim struct {
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
	Jti   string `json:"jti"`
	Name  string `json:"name"`
	Sub   string `json:"sub"`
	Iss   string `json:"iss"`
	Aud   string `json:"aud"`
}
