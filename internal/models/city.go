package models

type City struct {
	ID          uint   `gorm:"primaryKey"`
	DisplayName string `gorm:"column:display_name;index"`
	Code        string `gorm:"column:code;index"`

	Districts []District  `gorm:"foreignKey:CityID;references:ID;"`
	Aliases   []CityAlias `gorm:"foreignKey:CityID;references:ID;"`
}

func (City) TableName() string {
	return "cities"
}
