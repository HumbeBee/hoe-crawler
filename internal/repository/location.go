package repository

import (
	"errors"
	"github.com/HumbeBee/hoe-crawler/internal/models"
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
	var count int64
	err := r.db.Model(&models.District{}).Where("name = ? OR short_name = ? OR code = ? OR eng_name = ? OR display_name = ?", district, district, district, district, district).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("invalid location")
	}

	return nil
}

func (r *locationRepo) GetCityIDFromName(name string) (uint, error) {
	var city models.City
	err := r.db.Where("display_name = ? OR code = ?", name, name).First(&city).Error
	if err == nil {
		return city.ID, nil
	}

	var cityAlias models.CityAlias
	err = r.db.Where("alias = ?", name).First(&cityAlias).Error
	if err == nil {
		return cityAlias.CityID, nil
	}

	return city.ID, nil
}

func (r *locationRepo) GetDistrictIDFromName(name string) (uint, error) {
	var district models.District
	if err := r.db.Where("name = ?", name).Or("short_name = ?", name).Or("code = ?", name).Or("eng_name = ?", name).Or("display_name = ?", name).Or("eng_name = ?", name+" district").First(&district).Error; err != nil {
		return 0, err
	}

	return district.ID, nil
}
