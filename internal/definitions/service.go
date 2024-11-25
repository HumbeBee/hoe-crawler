package definitions

import (
	"github.com/haovoanh28/gai-webscraper/internal/dto"
	"github.com/haovoanh28/gai-webscraper/internal/models"
)

type MapperService interface {
	TransformHoe(rawInfo *dto.RawHoeData) *models.HoeInfo
}

type ValidateService interface {
	ValidateHoe(hoe *models.HoeInfo) error
}
