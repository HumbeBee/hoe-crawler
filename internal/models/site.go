package models

import "time"

type Site struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	BaseURL   string    `gorm:"column:base_url"`
	IsExpired bool      `gorm:"column:is_expired"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Site) TableName() string {
	return "sites"
}
