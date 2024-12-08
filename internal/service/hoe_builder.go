package service

import (
	"errors"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type hoeBuilder struct {
	hoeRepo            repository.HoeRepository
	workingHistoryRepo repository.WorkingHistoryRepository
	locationRepo       repository.LocationRepository
	logger             *logutil.Logger
	scraper            interfaces.Scraper
}

func NewHoeBuilder() *hoeBuilder {
	return &hoeBuilder{}
}

func (b *hoeBuilder) WithHoeRepo(hoeRepo repository.HoeRepository) *hoeBuilder {
	b.hoeRepo = hoeRepo
	return b
}

func (b *hoeBuilder) WithWorkingHistoryRepo(workingHistoryRepo repository.WorkingHistoryRepository) *hoeBuilder {
	b.workingHistoryRepo = workingHistoryRepo
	return b
}

func (b *hoeBuilder) WithLocationRepo(locationRepo repository.LocationRepository) *hoeBuilder {
	b.locationRepo = locationRepo
	return b
}

func (b *hoeBuilder) WithLogger(logger *logutil.Logger) *hoeBuilder {
	b.logger = logger
	return b
}

func (b *hoeBuilder) WithScraper(scraper interfaces.Scraper) *hoeBuilder {
	b.scraper = scraper
	return b
}

func (b *hoeBuilder) Build() (interfaces.HoeService, error) {
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

// ================================================================
func (b *hoeBuilder) validate() error {
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
