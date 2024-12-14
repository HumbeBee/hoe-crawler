package models

type City struct {
	ID          uint   `gorm:"primaryKey"`
	DisplayName string `gorm:"column:display_name"`
	Name        string `gorm:"column:name;index"`
	ShortName   string `gorm:"column:short_name;index"`
	EngName     string `gorm:"column:eng_name;index"`
	Code        string `gorm:"column:code;index"`

	Districts []District `gorm:"foreignKey:CityID;references:ID;"`
}

func (City) TableName() string {
	return "cities"
}
