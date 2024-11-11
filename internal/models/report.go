package models

type HoeReport struct {
	ID        uint   `gorm:"primaryKey"`
	HoeID     uint   `gorm:"column:hoe_id"`
	ReportURL string `gorm:"column:report_url"`
}

func (HoeReport) TableName() string {
	return "hoe_reports"
}
