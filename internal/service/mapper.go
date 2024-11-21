package service

import (
	"strings"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/dto"
	"github.com/haovoanh28/gai-webscraper/internal/models"
)

type MapperService interface {
	TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo
}

type mapperService struct{}

func NewMapperService() MapperService {
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

func (s *mapperService) mapStatus(status string) definitions.HoeStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "đang rảnh", "online", "hoạt động":
		return definitions.HoeStatusActive
	case "bận", "offline":
		return definitions.HoeStatusInactive
	default:
		return definitions.HoeStatusUnknown
	}
}

func (s *mapperService) parseAddress(address string) definitions.ParsedAddress {
	// Trim any leading/trailing whitespace
	address = strings.TrimSpace(address)

	// Split the address into parts
	parts := strings.Split(address, ",")

	parsed := definitions.ParsedAddress{
		Street:   "",
		District: "",
	}

	if len(parts) >= 1 {
		parsed.Street = strings.TrimSpace(parts[0])
	}

	if len(parts) >= 2 {
		// Handle potential district abbreviations
		district := strings.TrimSpace(parts[1])
		district = strings.ReplaceAll(district, "Q.", "Quận ")
		district = strings.ReplaceAll(district, "Quan ", "Quận ")
		parsed.District = district
	}

	return parsed
}
