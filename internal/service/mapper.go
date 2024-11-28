package service

import (
	"strings"

	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/models"
)

type mapperService struct{}

func NewMapperService() definitions.MapperService {
	return &mapperService{}
}

func (s *mapperService) TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo {
	rawInfo.Price = s.normalizePrice(rawInfo.Price)
	rawInfo.Phone = s.normalizePhone(rawInfo.Phone)

	return &models.HoeInfo{
		Name:      strings.TrimSpace(rawInfo.Name),
		Phone:     rawInfo.Phone,
		BirthYear: rawInfo.BirthYear,
		Height:    rawInfo.Height,
		Weight:    rawInfo.Weight,
		Country:   rawInfo.Country,

		Profiles: []models.HoeProfile{
			{
				SiteID:   rawInfo.SiteID,
				OriginID: rawInfo.OriginID,
				Url:      rawInfo.Url,
				ImageUrl: rawInfo.ImageUrl,
				Price:    rawInfo.Price,
				Area:     rawInfo.Area,
				Provider: rawInfo.Provider,
				Status:   s.mapStatus(rawInfo.Status),
				Service:  rawInfo.Service,
				Duration: rawInfo.Duration,
				WorkTime: rawInfo.WorkTime,
			},
		},
	}
}

func (s *mapperService) normalizePrice(price string) string {
	// First clean up any special characters and spaces
	price = strings.ReplaceAll(price, "\u00A0", "")
	price = strings.ReplaceAll(price, ",", "")
	price = strings.ToLower(price)

	// Remove any spaces
	price = strings.ReplaceAll(price, " ", "")

	// Remove 'k' and 'vnd'
	price = strings.ReplaceAll(price, "k", "")
	price = strings.ReplaceAll(price, "vnd", "")

	return price
}

func (s *mapperService) normalizePhone(phone string) string {
	phone = strings.ReplaceAll(phone, ".", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, " ", "")

	return phone
}

func (s *mapperService) mapStatus(status string) models.HoeStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "đang rảnh", "online", "hoạt động":
		return models.HoeStatusActive
	case "bận", "offline":
		return models.HoeStatusInactive
	default:
		return models.HoeStatusUnknown
	}
}
