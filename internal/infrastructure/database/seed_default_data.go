package database

import (
	"fmt"

	"github.com/HumbeBee/hoe-crawler/internal/models"
	"gorm.io/gorm"
)

func SeedDefaultData(db *gorm.DB) error {
	if err := seedSites(db); err != nil {
		return err
	}

	if err := seedCities(db); err != nil {
		return err
	}

	if err := seedFailedURLs(db); err != nil {
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
			Name:        "Hồ Chí Minh",
			Code:        "HCM",
			DisplayName: "Thành phố Hồ Chí Minh",
			EngName:     "Ho Chi Minh City",
			Districts: []models.District{
				{Name: "Quận 1", ShortName: "Q1", Code: "Q1", DisplayName: "Quận 1", EngName: "District 1"},
				{Name: "Quận 2", ShortName: "Q2", Code: "Q2", DisplayName: "Quận 2", EngName: "District 2"},
				{Name: "Quận 3", ShortName: "Q3", Code: "Q3", DisplayName: "Quận 3", EngName: "District 3"},
				{Name: "Quận 4", ShortName: "Q4", Code: "Q4", DisplayName: "Quận 4", EngName: "District 4"},
				{Name: "Quận 5", ShortName: "Q5", Code: "Q5", DisplayName: "Quận 5", EngName: "District 5"},
				{Name: "Quận 6", ShortName: "Q6", Code: "Q6", DisplayName: "Quận 6", EngName: "District 6"},
				{Name: "Quận 7", ShortName: "Q7", Code: "Q7", DisplayName: "Quận 7", EngName: "District 7"},
				{Name: "Quận 8", ShortName: "Q8", Code: "Q8", DisplayName: "Quận 8", EngName: "District 8"},
				{Name: "Quận 9", ShortName: "Q9", Code: "Q9", DisplayName: "Quận 9", EngName: "District 9"},
				{Name: "Quận 10", ShortName: "Q10", Code: "Q10", DisplayName: "Quận 10", EngName: "District 10"},
				{Name: "Quận 11", ShortName: "Q11", Code: "Q11", DisplayName: "Quận 11", EngName: "District 11"},
				{Name: "Quận 12", ShortName: "Q12", Code: "Q12", DisplayName: "Quận 12", EngName: "District 12"},
				{Name: "Quận Bình Tân", ShortName: "BTN", Code: "BTN", DisplayName: "Quận Bình Tân", EngName: "Binh Tan District"},
				{Name: "Quận Bình Thạnh", ShortName: "BTH", Code: "BTH", DisplayName: "Quận Bình Thạnh", EngName: "Binh Thanh District"},
				{Name: "Quận Gò Vấp", ShortName: "GV", Code: "GV", DisplayName: "Quận Gò Vấp", EngName: "Go Vap District"},
				{Name: "Quận Phú Nhuận", ShortName: "PN", Code: "PN", DisplayName: "Quận Phú Nhuận", EngName: "Phu Nhuan District"},
				{Name: "Quận Tân Bình", ShortName: "TB", Code: "TB", DisplayName: "Quận Tân Bình", EngName: "Tan Binh District"},
				{Name: "Quận Tân Phú", ShortName: "TP", Code: "TP", DisplayName: "Quận Tân Phú", EngName: "Tan Phu District"},
				{Name: "Thành phố Thủ Đức", ShortName: "TD", Code: "TD", DisplayName: "Thành phố Thủ Đức", EngName: "Thu Duc City"},
				{Name: "Huyện Bình Chánh", ShortName: "BCH", Code: "BCH", DisplayName: "Huyện Bình Chánh", EngName: "Binh Chanh District"},
				{Name: "Huyện Cần Giờ", ShortName: "CG", Code: "CG", DisplayName: "Huyện Cần Giờ", EngName: "Can Gio District"},
				{Name: "Huyện Củ Chi", ShortName: "CC", Code: "CC", DisplayName: "Huyện Củ Chi", EngName: "Cu Chi District"},
				{Name: "Huyện Hóc Môn", ShortName: "HM", Code: "HM", DisplayName: "Huyện Hóc Môn", EngName: "Hoc Mon District"},
				{Name: "Huyện Nhà Bè", ShortName: "NB", Code: "NB", DisplayName: "Huyện Nhà Bè", EngName: "Nha Be District"},
			},
		},
		{
			Name:        "Hà Nội",
			Code:        "HN",
			DisplayName: "Thành phố Hà Nội",
			EngName:     "Hanoi",
			Districts: []models.District{
				{Name: "Ba Đình", ShortName: "BD", Code: "BD", DisplayName: "Quận Ba Đình", EngName: "Ba Dinh District"},
				{Name: "Hoàn Kiếm", ShortName: "HK", Code: "HK", DisplayName: "Quận Hoàn Kiếm", EngName: "Hoan Kiem District"},
				{Name: "Tây Hồ", ShortName: "TH", Code: "TH", DisplayName: "Quận Tây Hồ", EngName: "Tay Ho District"},
				{Name: "Long Biên", ShortName: "LB", Code: "LB", DisplayName: "Quận Long Biên", EngName: "Long Bien District"},
				{Name: "Cầu Giấy", ShortName: "CG", Code: "CG", DisplayName: "Quận Cầu Giấy", EngName: "Cau Giay District"},
				{Name: "Đống Đa", ShortName: "DD", Code: "DD", DisplayName: "Quận Đống Đa", EngName: "Dong Da District"},
				{Name: "Hai Bà Trưng", ShortName: "HBT", Code: "HBT", DisplayName: "Quận Hai Bà Trưng", EngName: "Hai Ba Trung District"},
				{Name: "Hoàng Mai", ShortName: "HM", Code: "HM", DisplayName: "Quận Hoàng Mai", EngName: "Hoang Mai District"},
				{Name: "Thanh Xuân", ShortName: "TX", Code: "TX", DisplayName: "Quận Thanh Xuân", EngName: "Thanh Xuan District"},
				{Name: "Sóc Sơn", ShortName: "SS", Code: "SS", DisplayName: "Huyện Sóc Sơn", EngName: "Soc Son District"},
				{Name: "Đông Anh", ShortName: "DA", Code: "DA", DisplayName: "Huyện Đông Anh", EngName: "Dong Anh District"},
				{Name: "Gia Lâm", ShortName: "GL", Code: "GL", DisplayName: "Huyện Gia Lâm", EngName: "Gia Lam District"},
				{Name: "Nam Từ Liêm", ShortName: "NTL", Code: "NTL", DisplayName: "Quận Nam Từ Liêm", EngName: "Nam Tu Liem District"},
				{Name: "Thanh Trì", ShortName: "TT", Code: "TT", DisplayName: "Huyện Thanh Trì", EngName: "Thanh Tri District"},
				{Name: "Bắc Từ Liêm", ShortName: "BTL", Code: "BTL", DisplayName: "Quận Bắc Từ Liêm", EngName: "Bac Tu Liem District"},
				{Name: "Mê Linh", ShortName: "ML", Code: "ML", DisplayName: "Huyện Mê Linh", EngName: "Me Linh District"},
				{Name: "Hà Đông", ShortName: "HD", Code: "HD", DisplayName: "Quận Hà Đông", EngName: "Ha Dong District"},
				{Name: "Sơn Tây", ShortName: "ST", Code: "ST", DisplayName: "Thị xã Sơn Tây", EngName: "Son Tay Town"},
				{Name: "Ba Vì", ShortName: "BV", Code: "BV", DisplayName: "Huyện Ba Vì", EngName: "Ba Vi District"},
				{Name: "Phúc Thọ", ShortName: "PT", Code: "PT", DisplayName: "Huyện Phúc Thọ", EngName: "Phuc Tho District"},
				{Name: "Đan Phượng", ShortName: "DP", Code: "DP", DisplayName: "Huyện Đan Phượng", EngName: "Dan Phuong District"},
				{Name: "Hoài Đức", ShortName: "HD", Code: "HD", DisplayName: "Huyện Hoài Đức", EngName: "Hoai Duc District"},
				{Name: "Quốc Oai", ShortName: "QO", Code: "QO", DisplayName: "Huyện Quốc Oai", EngName: "Quoc Oai District"},
				{Name: "Thạch Thất", ShortName: "TT", Code: "TT", DisplayName: "Huyện Thạch Thất", EngName: "Thach That District"},
				{Name: "Chương Mỹ", ShortName: "CM", Code: "CM", DisplayName: "Huyện Chương Mỹ", EngName: "Chuong My District"},
				{Name: "Thanh Oai", ShortName: "TO", Code: "TO", DisplayName: "Huyện Thanh Oai", EngName: "Thanh Oai District"},
				{Name: "Thường Tín", ShortName: "TT", Code: "TT", DisplayName: "Huyện Thường Tín", EngName: "Thuong Tin District"},
				{Name: "Phú Xuyên", ShortName: "PX", Code: "PX", DisplayName: "Huyện Phú Xuyên", EngName: "Phu Xuyen District"},
				{Name: "Ứng Hòa", ShortName: "UH", Code: "UH", DisplayName: "Huyện Ứng Hòa", EngName: "Ung Hoa District"},
				{Name: "Mỹ Đức", ShortName: "MD", Code: "MD", DisplayName: "Huyện Mỹ Đức", EngName: "My Duc District"},
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
			for _, province := range city.Districts {
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

func seedFailedURLs(db *gorm.DB) error {
	failedURLs := []models.FailedURL{
		{URL: "/gai-goi/chi-tiet/61543/reup-mie-ly-cocktail-girl-cua-nhung-bua-tiec-sex", Type: "detail", SiteID: 1, LastError: "test error", RetryCount: 0},
		{URL: "/gai-goi/chi-tiet/62010/reup-be-moc-tra-dam-mup-cute-teen-xinh-sang-body-nuot-thom", Type: "detail", SiteID: 1, LastError: "test error", RetryCount: 0},
		{URL: "/gai-goi/chi-tiet/61614/rep-upbao-han-tretrung-xinh-xan-lam-tinh-chuyen-nghiep", Type: "detail", SiteID: 1, LastError: "test error", RetryCount: 0},
		{URL: "/gai-goi/chi-tiet/63804/gai-xinh-my-linh-da-trang-mat-xinh-3-vong-sieu-dep-luon", Type: "detail", SiteID: 1, LastError: "test error", RetryCount: 0},
	}

	result := db.Create(&failedURLs).Error
	if result != nil {
		return result
	}

	return nil
}
