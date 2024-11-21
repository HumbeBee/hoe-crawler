package service

import (
	"errors"
	"github.com/haovoanh28/gai-webscraper/internal/definitions"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"strings"
)

type ValidateService interface {
}

type validateService struct {
	locationRepo repository.LocationRepository
}

func NewValidateService(locationRepo repository.LocationRepository) ValidateService {
	return &validateService{locationRepo: locationRepo}
}

func (s *validateService) ValidateHoe(hoe *models.HoeInfo) error {
	return nil
}

func (s *validateService) ValidateLocation(location string) error {
	if location == "" {
		return errors.New("location cannot be empty")
	}

	return nil
}

func (s *validateService) parseAddress(address string) definitions.ParsedAddress {
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
