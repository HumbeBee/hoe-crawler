package service

import (
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/models"

	"strings"
)

type validateService struct {
}

func NewValidateService() definitions.ValidateService {
	return &validateService{}
}

func (s *validateService) ValidateHoe(hoe *models.HoeInfo) error {

	return nil
}

func parseLocation(location string) definitions.ParsedAddress {
	// Trim any leading/trailing whitespace
	location = strings.TrimSpace(location)

	// Split the address into parts
	parts := strings.Split(location, ",")

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
