package models

type District struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Code   string `gorm:"column:code"`
	CityID uint   `gorm:"column:city_id"`
}

func (District) TableName() string {
	return "districts"
}
