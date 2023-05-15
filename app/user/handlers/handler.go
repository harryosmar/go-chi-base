package handlers

import (
	"github.com/harryosmar/go-chi-base/app/user/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}
