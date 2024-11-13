package service

import (
	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type HoeService struct {
	repo   repository.HoeRepository
	logger *logutil.Logger
}

func NewHoeService(repo repository.HoeRepository, logger *logutil.Logger) *HoeService {
	return &HoeService{repo: repo, logger: logger}
}

func (s *HoeService) ProcessHoe(hoe *models.HoeInfo) error {
	s.logger.Info("Processing hoe: ", hoe.Name)
	return nil
}
