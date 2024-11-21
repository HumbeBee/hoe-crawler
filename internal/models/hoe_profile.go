package models

import (
	"time"

	"github.com/haovoanh28/gai-webscraper/internal/definitions"
)

type HoeProfile struct {
	ID        uint                  `gorm:"primaryKey"`
	HoeID     uint                  `gorm:"column:hoe_id"`
	SiteID    uint                  `gorm:"column:site_id;index"`
	OriginID  string                `gorm:"column:origin_id"`
	Url       string                `gorm:"column:url"`
	ImageUrl  string                `gorm:"column:image_url"`
	Price     string                `gorm:"column:price"`
	Provider  string                `gorm:"column:provider"`
	Status    definitions.HoeStatus `gorm:"column:status;type:enum('active','inactive','unknown')"`
	Service   string                `gorm:"column:service"`
	Duration  string                `gorm:"column:duration"`
	WorkTime  string                `gorm:"column:work_time"`
	CreatedAt time.Time             `gorm:"column:created_at"`
	UpdatedAt time.Time             `gorm:"column:updated_at"`

	Hoe     HoeInfo     `gorm:"foreignKey:HoeID"`
	Site    Site        `gorm:"foreignKey:SiteID"`
	Reports []HoeReport `gorm:"foreignKey:HoeProfileID"`
}

func (p *HoeProfile) TableName() string {
	return "hoe_profiles"
}
