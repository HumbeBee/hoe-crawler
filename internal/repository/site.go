package repository

import (
	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

type SiteRepository interface {
	GetSiteByName(name string) (*models.Site, error)
	GetSiteByID(id uint) (*models.Site, error)
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

func (r *siteRepo) GetSiteByID(id uint) (*models.Site, error) {
	var site models.Site
	if err := r.db.Model(&models.Site{}).Where("id = ?", id).First(&site).Error; err != nil {
		return nil, err
	}

	return &site, nil
}
