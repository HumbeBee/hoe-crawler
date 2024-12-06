package service

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/models"

	"strings"
)

type validateService struct {
}

func NewValidateService() interfaces.ValidateService {
	return &validateService{}
}

func (s *validateService) ValidateHoe(hoe *models.HoeInfo) error {
	if hoe.Phone == "" {
		return fmt.Errorf("phone number is required")
	}

	if len(hoe.Profiles) == 0 {
		return fmt.Errorf("hoe must have at least one profile")
	}

	profile := hoe.Profiles[0]
	if profile.SiteID == 0 {
		return fmt.Errorf("profile must have site ID")
	}

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
