package repository

import (
	"github.com/haovoanh28/gai-webscraper/internal/db"
	"github.com/haovoanh28/gai-webscraper/internal/models"
)

type HoeRepository interface {
	Save(hoe *models.HoeInfo) error
}

type hoeRepo struct {
	dbo *db.DBO
}

func NewHoeRepository(dbo *db.DBO) HoeRepository {
	return &hoeRepo{dbo: dbo}
}

func (r *hoeRepo) Save(hoe *models.HoeInfo) error {
	return r.dbo.InsertHoe(hoe)
}
