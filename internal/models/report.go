package models

import "time"

type HoeReport struct {
	ID        uint      `gorm:"primaryKey"`
	HoeID     uint      `gorm:"column:hoe_id"`
	ReportURL string    `gorm:"column:report_url"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (HoeReport) TableName() string {
	return "hoe_reports"
}
