package interfaces

import (
	"github.com/HumbeBee/hoe-crawler/internal/dto"
	"github.com/HumbeBee/hoe-crawler/internal/models"
)

type MapperService interface {
	TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo
}

type ValidateService interface {
	ValidateHoe(hoe *models.HoeInfo) error
}
