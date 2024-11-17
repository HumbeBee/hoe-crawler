package models

type City struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"column:name"`
	Code      string     `gorm:"column:code"`
	Provinces []Province `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
}

func (City) TableName() string {
	return "cities"
}
