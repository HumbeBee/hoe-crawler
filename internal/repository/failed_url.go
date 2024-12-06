package repository

import (
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

type FailedURLRepository interface {
	GetFailedURLs() ([]string, error)
	Save(failedURL *models.FailedURL) error
}

type failedURLRepo struct {
	db *gorm.DB
}

func NewFailedURLRepository(db *gorm.DB) FailedURLRepository {
	return &failedURLRepo{db: db}
}

func (fu *failedURLRepo) GetFailedURLs() ([]string, error) {
	const LIMIT = 20

	var failedURLs []string
	if err := fu.db.Limit(LIMIT).Find(&failedURLs).Error; err != nil {
		return nil, err
	}

	return failedURLs, nil
}

func (fu *failedURLRepo) Save(failedURL *models.FailedURL) error {
	tx := fu.db.Begin()

	if err := tx.Create(failedURL).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
