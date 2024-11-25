package models

type City struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"column:name;index"`
	Code      string     `gorm:"column:code;index"`
	Districts []District `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:CASCADE"`
}

func (City) TableName() string {
	return "cities"
}
