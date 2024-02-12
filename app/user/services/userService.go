package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/harryosmar/go-chi-base/app/user/entities"
	"github.com/harryosmar/go-chi-base/app/user/repositories"
	"github.com/harryosmar/go-chi-base/core/errors"
	"github.com/harryosmar/go-chi-base/core/logger"
	"gorm.io/gorm"
	"time"
)

//go:generate mockgen -destination=mocks/mock_UserService.go -package=mocks . UserService
type UserService interface {
	Authenticate(ctx context.Context, c entities.AuthenticateRequest) (*entities.AuthenticateResponse, error)
	ValidateToken(ctx context.Context, c entities.ValidateTokenRequest) (*entities.JwtClaim, error)
}

type userService struct {
	tokenIssuer       string
	accountRepository repositories.AccountRepository
}

func NewUserService(tokenIssuer string, accountRepository repositories.AccountRepository) UserService {
	return &userService{tokenIssuer: tokenIssuer, accountRepository: accountRepository}
}

func (u userService) Authenticate(ctx context.Context, c entities.AuthenticateRequest) (*entities.AuthenticateResponse, error) {
	logEntry := logger.GetLogEntryFromCtx(ctx)
	account, err := u.accountRepository.GetProfileByUsername(ctx, c.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, codes.ErrLoginCredential
		}

		logEntry.Error(err)
		return nil, err
	}

	_ = entities.JwtClaim{
		Email: account.Email,
		Exp:   time.Now().Unix(),
		Iat:   time.Now().Unix() + 86400,
		Jti:   uuid.New().String(),
		Name:  account.FullName,
		Sub:   account.Username,
		Iss:   "api.kursus-masak.id",
		Aud:   "www.kursus-masak.id",
	}
	return &entities.AuthenticateResponse{Token: ""}, nil
}

func (u userService) ValidateToken(ctx context.Context, c entities.ValidateTokenRequest) (*entities.JwtClaim, error) {
	return &entities.JwtClaim{}, nil
}
