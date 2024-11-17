package models

type Province struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"column:name"`
	Code   string `gorm:"column:code"`
	CityID uint   `gorm:"column:city_id"`
}

func (Province) TableName() string {
	return "provinces"
}
