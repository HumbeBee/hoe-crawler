package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

type SiteRepository interface {
	GetSiteByName(name string) (*models.Site, error)
}

type siteRepo struct {
	db *gorm.DB
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepo{db: db}
}

func (r *siteRepo) GetSiteByName(name string) (*models.Site, error) {
	var site models.Site
	if err := r.db.Model(&models.Site{}).Where("name = ?", name).First(&site).Error; err != nil {
		return nil, err
	}

	return &site, nil
}
