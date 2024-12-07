package repository

import (
	"errors"
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

type FailedURLRepository interface {
	GetFailedURLs() ([]string, error)
	FindFailedURL(url string, sideID uint) (*models.FailedURL, error)
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
	return fu.db.Create(failedURL).Error
}

func (fu *failedURLRepo) FindFailedURL(url string, siteID uint) (*models.FailedURL, error) {
	var failedURL models.FailedURL
	err := fu.db.Where("url = ? AND site_id = ?", url, siteID).First(&failedURL).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &failedURL, nil
}
