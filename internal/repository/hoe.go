package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/infrastructure/database"
	"github.com/haovoanh28/gai-webscraper/internal/models"
)

type HoeRepository interface {
	Save(hoe *models.HoeInfo) error
}

type hoeRepo struct {
	dbo *database.DBO
}

func NewHoeRepository(dbo *database.DBO) HoeRepository {
	return &hoeRepo{dbo: dbo}
}

func (r *hoeRepo) Save(hoe *models.HoeInfo) error {
	return r.dbo.InsertHoe(hoe)
}
