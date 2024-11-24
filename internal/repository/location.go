package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

type LocationRepository interface {
	CheckValidLocation(district string) error
}

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepo{db: db}
}

func (r *locationRepo) CheckValidLocation(district string) error {
	if err := r.db.Where("name = ? ", district).Or("short_name = ?").Or("code = ? ").First(&models.District{}).Error; err != nil {
		return err
	}

	return nil
}
