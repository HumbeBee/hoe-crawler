package service

import (
	"errors"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/browser"
	"github.com/HumbeBee/hoe-crawler/internal/interfaces"
	"github.com/HumbeBee/hoe-crawler/internal/repository"
	"github.com/HumbeBee/hoe-crawler/internal/utils/logutil"
)

type failedURLBuilder struct {
	siteID             uint
	logger             *logutil.Logger
	browserRateLimiter *browser.RateLimiter
	failedURLRepo      repository.FailedURLRepository
	siteRepo           repository.SiteRepository
}

func NewFailedURLBuilder() *failedURLBuilder {
	return &failedURLBuilder{}
}

func (b *failedURLBuilder) WithSiteID(siteID uint) *failedURLBuilder {
	b.siteID = siteID
	return b
}

func (b *failedURLBuilder) WithFailedURLRepo(failedURLRepo repository.FailedURLRepository) *failedURLBuilder {
	b.failedURLRepo = failedURLRepo
	return b
}

func (b *failedURLBuilder) WithLogger(logger *logutil.Logger) *failedURLBuilder {
	b.logger = logger
	return b
}

func (b *failedURLBuilder) WithSiteRepo(siteRepo repository.SiteRepository) *failedURLBuilder {
	b.siteRepo = siteRepo
	return b
}

func (b *failedURLBuilder) WithBrowserRateLimiter(browserRateLimiter *browser.RateLimiter) *failedURLBuilder {
	b.browserRateLimiter = browserRateLimiter
	return b
}

func (b *failedURLBuilder) Build() (interfaces.FailedURLService, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}

	return &failedURLService{
		siteID:             b.siteID,
		failedURLRepo:      b.failedURLRepo,
		logger:             b.logger,
		siteRepo:           b.siteRepo,
		browserRateLimiter: b.browserRateLimiter,
	}, nil
}

// ================================================================
func (b *failedURLBuilder) validate() error {
	if b.siteID == 0 {
		return errors.New("siteID is required")
	}
	if b.browserRateLimiter == nil {
		return errors.New("browserRateLimiter is required")
	}
	if b.failedURLRepo == nil {
		return errors.New("failedURLRepo is required")
	}
	if b.siteRepo == nil {
		return errors.New("siteRepo is required")
	}

	// Set defaults for optional stuff
	if b.logger == nil {
		b.logger = logutil.NewLogger(logutil.INFO)
	}

	return nil
}
