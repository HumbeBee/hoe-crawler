package repository

import "gorm.io/gorm"

type LocationRepository interface {
}

type locationRepo struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepo{db: db}
}
