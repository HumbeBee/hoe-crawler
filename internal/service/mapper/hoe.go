package mapper

import (
	"strings"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/dto"
	"github.com/haovoanh28/gai-webscraper/internal/models"
)

func TransformHoe(rawInfo dto.RawHoeData) *models.HoeInfo {
	rawInfo.Price = normalizePrice(rawInfo.Price)
	rawInfo.Phone = normalizePhone(rawInfo.Phone)

	return &models.HoeInfo{
		Name:      strings.TrimSpace(rawInfo.Name),
		Phone:     normalizePhone(rawInfo.Phone),
		BirthYear: rawInfo.BirthYear,
		Height:    rawInfo.Height,
		Weight:    rawInfo.Weight,
		Country:   rawInfo.Country,

		Profiles: []models.HoeProfile{
			{
				SiteID: rawInfo.SiteID,
				Status: mapStatus(rawInfo.Status),
			},
		},
	}
}

func normalizePrice(price string) string {
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

func normalizePhone(phone string) string {
	phone = strings.ReplaceAll(phone, ".", "")
	return phone
}

func mapStatus(status string) definitions.HoeStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "đang rảnh", "online":
		return definitions.HoeStatusActive
	case "bận", "offline":
		return definitions.HoeStatusInactive
	default:
		return definitions.HoeStatusUnknown
	}
}
