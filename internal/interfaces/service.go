package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/models"
)

type HoeService interface {
	ProcessListPage() error
	ProcessDetailPage(url string) error
}

type MapperService interface {
	TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo
}

type ValidateService interface {
	ValidateHoe(hoe *models.HoeInfo) error
}

type FailedURLService interface {
	TrackFailedURL(url string, err error)
	RetryFailedURLs() error
}
