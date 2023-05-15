package repositories

import (
	"context"
	"gorm.io/gorm"
)

type CredentialRepository interface {
	GetUserByCredential(ctx context.Context, username string, password string) (int64, error)
}

type credentialRepository struct {
	db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) CredentialRepository {
	return &credentialRepository{db: db}
}

func (u credentialRepository) GetUserByCredential(ctx context.Context, username string, password string) (int64, error) {
	if username == "febri@angel.com" {
		return 1, nil
	}

	return 0, gorm.ErrRecordNotFound
}
