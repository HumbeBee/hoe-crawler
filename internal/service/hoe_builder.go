package service

import (
	"errors"

	"github.com/haovoanh28/gai-webscraper/internal/interfaces"
	"github.com/haovoanh28/gai-webscraper/internal/repository"
	"github.com/haovoanh28/gai-webscraper/internal/utils/logutil"
)

type HoeBuilder struct {
	hoeRepo            repository.HoeRepository
	workingHistoryRepo repository.WorkingHistoryRepository
	locationRepo       repository.LocationRepository
	logger             *logutil.Logger
	scraper            interfaces.Scraper
}

func NewHoeBuilder() *HoeBuilder {
	return &HoeBuilder{}
}

func (b *HoeBuilder) WithHoeRepo(hoeRepo repository.HoeRepository) *HoeBuilder {
	b.hoeRepo = hoeRepo
	return b
}

func (b *HoeBuilder) WithWorkingHistoryRepo(workingHistoryRepo repository.WorkingHistoryRepository) *HoeBuilder {
	b.workingHistoryRepo = workingHistoryRepo
	return b
}

func (b *HoeBuilder) WithLocationRepo(locationRepo repository.LocationRepository) *HoeBuilder {
	b.locationRepo = locationRepo
	return b
}

func (b *HoeBuilder) WithLogger(logger *logutil.Logger) *HoeBuilder {
	b.logger = logger
	return b
}

func (b *HoeBuilder) WithScraper(scraper interfaces.Scraper) *HoeBuilder {
	b.scraper = scraper
	return b
}

func (b *HoeBuilder) validate() error {
	if b.hoeRepo == nil {
		return errors.New("hoeRepo is required")
	}
	if b.workingHistoryRepo == nil {
		return errors.New("workingHistoryRepo is required")
	}
	if b.locationRepo == nil {
		return errors.New("locationRepo is required")
	}
	if b.scraper == nil {
		return errors.New("scraper is required")
	}

	// Set defaults for optional stuff
	if b.logger == nil {
		b.logger = logutil.NewLogger(logutil.INFO)
	}

	return nil
}

func (b *HoeBuilder) Build() (HoeService, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	return &hoeService{
		locationRepo:       b.locationRepo,
		hoeRepo:            b.hoeRepo,
		workingHistoryRepo: b.workingHistoryRepo,
		logger:             b.logger,
		scraper:            b.scraper,
		mapperService:      NewMapperService(),
		validateService:    NewValidateService(),
	}, nil
}
