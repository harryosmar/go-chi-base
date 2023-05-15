package repositories

import (
	"context"
	"github.com/harryosmar/go-chi-base/app/user/entities"
	"gorm.io/gorm"
	"time"
)

type ProfileRepository interface {
	GetProfileByUserId(ctx context.Context, userId int64) (entities.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (p profileRepository) GetProfileByUserId(ctx context.Context, userId int64) (entities.Profile, error) {
	return entities.Profile{
		Name: "Febri Angel",
		BirthDate: func() time.Time {
			parse, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00+07:00")
			return parse
		}(),
	}, nil
}
