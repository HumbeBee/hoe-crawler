package database

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

func SeedDefaultData(db *gorm.DB) error {
	if err := seedSites(db); err != nil {
		return err
	}

	if err := seedCities(db); err != nil {
		return err
	}

	return nil
}

func seedSites(db *gorm.DB) error {
	sites := []models.Site{
		{
			Name:      "gaito",
			BaseURL:   "https://gaito.so",
			IsExpired: false,
		},
		{
			Name:      "gaigu",
			BaseURL:   "https://gaigu31.tv",
			IsExpired: false,
		},
	}

	for _, site := range sites {
		result := db.Where(models.Site{Name: site.Name}).
			FirstOrCreate(&site)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func seedCities(db *gorm.DB) error {
	cities := []models.City{
		{
			Name: "Hồ Chí Minh",
			Code: "HCM",
			Provinces: []models.District{
				{Name: "Quận 1", Code: "Q1"},
				{Name: "Quận 2", Code: "Q2"},
				{Name: "Quận 3", Code: "Q3"},
				{Name: "Quận 4", Code: "Q4"},
				{Name: "Quận 5", Code: "Q5"},
				{Name: "Quận 6", Code: "Q6"},
				{Name: "Quận 7", Code: "Q7"},
				{Name: "Quận 8", Code: "Q8"},
				{Name: "Quận 9", Code: "Q9"},
				{Name: "Quận 10", Code: "Q10"},
				{Name: "Quận 11", Code: "Q11"},
				{Name: "Quận 12", Code: "Q12"},
				{Name: "Quận Bình Tân", Code: "BTN"},
				{Name: "Quận Bình Thạnh", Code: "BTH"},
				{Name: "Quận Bình Chánh", Code: "BCH"},
				{Name: "Quận Gò Vấp", Code: "GV"},
				{Name: "Quận Phú Nhuận", Code: "PN"},
				{Name: "Quận Tân Bình", Code: "TB"},
				{Name: "Quận Tân Phú", Code: "TP"},
				{Name: "Quận Thủ Đức", Code: "TD"},
			},
		},
		{
			Name: "Hà Nội",
			Code: "HN",
			Provinces: []models.District{
				{Name: "Quận Ba Đình", Code: "BD"},
				{Name: "Quận Hoàn Kiếm", Code: "HK"},
				{Name: "Quận Hai Bà Trưng", Code: "HBT"},
				{Name: "Quận Đống Đa", Code: "DD"},
				{Name: "Quận Tây Hồ", Code: "TH"},
				{Name: "Quận Cầu Giấy", Code: "CG"},
				{Name: "Quận Thanh Xuân", Code: "TX"},
				{Name: "Quận Hoàng Mai", Code: "HM"},
				{Name: "Quận Long Biên", Code: "LB"},
				{Name: "Quận Nam Từ Liêm", Code: "NTL"},
				{Name: "Quận Bắc Từ Liêm", Code: "BTL"},
				{Name: "Quận Hà Đông", Code: "HD"},
			},
		},
	}

	for _, city := range cities {
		var existingCity models.City
		result := db.Where(models.City{Name: city.Name}).First(&existingCity)

		if result.Error != nil {
			// City doesn't exist, create new city with provinces
			if err := db.Create(&city).Error; err != nil {
				return fmt.Errorf("failed to create city %s: %v", city.Name, err)
			}
		} else {
			// City exists, check and create missing provinces
			for _, province := range city.Provinces {
				province.CityID = existingCity.ID
				result := db.Where(models.District{
					Name:   province.Name,
					CityID: existingCity.ID,
				}).FirstOrCreate(&province)

				if result.Error != nil {
					return fmt.Errorf("failed to create province %s: %v", province.Name, result.Error)
				}
			}
		}
	}

	return nil
}
