package services

import (
	"context"
	"github.com/harryosmar/go-chi-base/app/user/entities"
	"github.com/harryosmar/go-chi-base/app/user/repositories"
	codes "github.com/harryosmar/go-chi-base/errors"
	"github.com/harryosmar/go-chi-base/logger"
	"gorm.io/gorm"
)

type UserService interface {
	ValidateCredentials(ctx context.Context, c entities.ValidateCredentialRequest) (entities.ValidateCredentialResponse, error)
}

type userService struct {
	userRepo          repositories.CredentialRepository
	profileRepository repositories.ProfileRepository
}

func NewUserService(userRepo repositories.CredentialRepository, profileRepository repositories.ProfileRepository) UserService {
	return &userService{userRepo: userRepo, profileRepository: profileRepository}
}

func (u userService) ValidateCredentials(ctx context.Context, c entities.ValidateCredentialRequest) (entities.ValidateCredentialResponse, error) {
	logEntry := logger.GetLogEntryFromCtx(ctx)
	userId, err := u.userRepo.GetUserByCredential(ctx, c.Username, c.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.ValidateCredentialResponse{}, codes.ErrLoginCredential
		}

		logEntry.Error(err)
		return entities.ValidateCredentialResponse{}, err
	}

	profile, err := u.profileRepository.GetProfileByUserId(ctx, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.ValidateCredentialResponse{}, codes.ErrProfileNotExist
		}

		return entities.ValidateCredentialResponse{}, err
	}

	return entities.ValidateCredentialResponse{Message: "login berhasil", Profile: profile}, nil
}
