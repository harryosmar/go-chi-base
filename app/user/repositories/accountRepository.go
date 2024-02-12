package repositories

import (
	"context"
	"github.com/harryosmar/go-chi-base/app/user/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetProfileByUsername(ctx context.Context, username string) (*entities.Account, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) AccountRepository {
	return &profileRepository{db: db}
}

func (p profileRepository) GetProfileByUsername(ctx context.Context, username string) (*entities.Account, error) {
	return &entities.Account{
		Id:       1,
		Username: username,
		Email:    "febri@angel.com",
		FullName: "Febri & angel",
		Password: "123456",
	}, nil
}
