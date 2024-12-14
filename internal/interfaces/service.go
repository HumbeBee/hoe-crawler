package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/models"
)

type HoeService interface {
	ProcessListPage(baseURL string, relativeURL string) error
	ProcessDetailPage(baseURL string, relativeURL string) error
	GetRawHoeData(baseURL string, relativeURL string) (*dto.RawHoeData, error)
	ProcessRawHoeData(rawHoeData *dto.RawHoeData) error
}

type MapperService interface {
	TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo
}

type ValidateService interface {
	ValidateHoe(hoe *models.HoeInfo) error
}

type FailedURLService interface {
	TrackFailedURL(failedType models.FailedType, url string, err error)
	RetryFailedURLs() error
}
