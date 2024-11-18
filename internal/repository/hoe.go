package repository

import (
	"fmt"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"gorm.io/gorm"
)

type HoeRepository interface {
	Save(hoe *models.HoeInfo) error
}

type hoeRepo struct {
	db *gorm.DB
}

func NewHoeRepository(db *gorm.DB) HoeRepository {
	return &hoeRepo{db: db}
}

func (r *hoeRepo) Save(hoe *models.HoeInfo) error {
	transaction := r.db.Begin()

	if err := transaction.Create(hoe).Error; err != nil {
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		transaction.Rollback()
		return err
	}

	fmt.Println("Inserted")
	hoe.Print()

	return nil
}
