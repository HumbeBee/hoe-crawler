package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

type LocationRepository interface {
	CheckValidLocation(district string) error
	GetCityIDFromName(name string) (uint, error)
	GetDistrictIDFromName(name string) (uint, error)
}

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepo{db: db}
}

func (r *locationRepo) CheckValidLocation(district string) error {
	if err := r.db.Where("name = ?", district).Or("short_name = ?", district).Or("code = ? ", district).First(&models.District{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *locationRepo) GetCityIDFromName(name string) (uint, error) {
	var city models.City
	if err := r.db.Where("name = ?", name).Or("short_name = ?", name).Or("code = ?", name).Or("eng_name = ?", name).Or("display_name = ?", name).Or("eng_name = ?", name+" city").First(&city).Error; err != nil {
		return 0, err
	}

	return city.ID, nil
}

func (r *locationRepo) GetDistrictIDFromName(name string) (uint, error) {
	var district models.District
	if err := r.db.Where("name = ?", name).First(&district).Error; err != nil {
		return 0, err
	}

	return district.ID, nil
}
