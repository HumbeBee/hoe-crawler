package repository

import (
	"errors"
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

type FailedURLRepository interface {
	GetFailedURLs() ([]*models.FailedURL, error)
	FindFailedURL(url string, sideID uint) (*models.FailedURL, error)
	Save(failedURL *models.FailedURL) error
	Delete(failedURL *models.FailedURL) error
}

type failedURLRepo struct {
	db *gorm.DB
}

func NewFailedURLRepository(db *gorm.DB) FailedURLRepository {
	return &failedURLRepo{db: db}
}

func (fu *failedURLRepo) GetFailedURLs() ([]*models.FailedURL, error) {
	const LIMIT = 20

	var failedURLs []*models.FailedURL
	if err := fu.db.Limit(LIMIT).Find(&failedURLs).Error; err != nil {
		return nil, err
	}

	return failedURLs, nil
}

func (fu *failedURLRepo) Save(failedURL *models.FailedURL) error {
	return fu.db.Save(failedURL).Error
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

func (fu *failedURLRepo) Delete(failedURL *models.FailedURL) error {
	return fu.db.Delete(failedURL).Error
}
