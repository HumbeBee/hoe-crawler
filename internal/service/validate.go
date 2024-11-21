package service

import (
	"errors"
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
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
