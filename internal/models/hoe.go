package models

import (
	"fmt"
	"time"
)

type HoeInfo struct {
	ID        uint         `gorm:"primaryKey"`
	OriginID  string       `gorm:"column:origin_id;unique"`
	Url       string       `gorm:"column:url"`
	Name      string       `gorm:"column:name"`
	ImageUrl  string       `gorm:"column:image_url"`
	Price     string       `gorm:"column:price"`
	Phone     string       `gorm:"column:phone"`
	Address   string       `gorm:"column:address"`
	Provider  string       `gorm:"column:provider"`
	Status    string       `gorm:"column:status"`
	BirthYear string       `gorm:"column:birth_year"`
	Height    string       `gorm:"column:height"`
	Weight    string       `gorm:"column:weight"`
	Country   string       `gorm:"column:country"`
	Service   string       `gorm:"column:service"`
	Duration  string       `gorm:"column:duration"`
	WorkTime  string       `gorm:"column:work_time"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	Reports   []*HoeReport `gorm:"foreignKey:HoeID;references:ID"`
}

func (HoeInfo) TableName() string {
	return "hoes"
}

func (hoe *HoeInfo) Print() {
	fmt.Printf("=========== Hoe %s ===========\n", hoe.OriginID)
	fmt.Printf("Url: %s\n", hoe.Url)

	// Print all fields in Main info
	fmt.Printf("Name: %s\n", hoe.Name)
	fmt.Printf("Image url: %s\n", hoe.ImageUrl)
	fmt.Printf("Price: %s\n", hoe.Price)
	fmt.Printf("Phone: %s\n", hoe.Phone)
	fmt.Printf("Address: %s\n", hoe.Address)
	fmt.Printf("Provider: %s\n", hoe.Provider)
	fmt.Printf("Status: %s\n", hoe.Status)
	fmt.Printf("Birth year: %s\n", hoe.BirthYear)
	fmt.Printf("Height: %s\n", hoe.Height)
	fmt.Printf("Weight: %s\n", hoe.Weight)
	fmt.Printf("From: %s\n", hoe.Country)
	fmt.Printf("Service: %s\n", hoe.Service)
	fmt.Printf("Work Time: %s\n", hoe.WorkTime)
	fmt.Printf("Duration: %s\n", hoe.Duration)
	fmt.Print("Report count: ", len(hoe.Reports))

	fmt.Println("==============================")
	fmt.Print("\n\n")
}
