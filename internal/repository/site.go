package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

type SiteRepository interface {
	GetSiteIDFromName(name string) (int, error)
}

type siteRepo struct {
	db *gorm.DB
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepo{db: db}
}

func (r *siteRepo) GetSiteIDFromName(name string) (int, error) {
	var id int
	if err := r.db.Model(&models.Site{}).Where("name = ?", name).Select("id").Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}
