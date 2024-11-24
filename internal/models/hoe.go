package models

import (
	"fmt"
	"time"
)

type HoeInfo struct {
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"column:name"`
	Phone         string    `gorm:"column:phone;index;unique"`
	BirthYear     string    `gorm:"column:birth_year"`
	Height        string    `gorm:"column:height"`
	Weight        string    `gorm:"column:weight"`
	Country       string    `gorm:"column:country"`
	LastScrapedAt time.Time `gorm:"column:last_scraped_at"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`

	// Site-specific profiles
	Profiles []HoeProfile `gorm:"foreignKey:HoeID;references:ID"`
	// Working histories
	WorkingHistories []WorkingHistory `gorm:"foreignKey:HoeID;references:ID"`
}

func (*HoeInfo) TableName() string {
	return "hoes"
}

func (h *HoeInfo) GetProfileBySite(siteID uint) *HoeProfile {
	for i := range h.Profiles {
		if h.Profiles[i].SiteID == siteID {
			return &h.Profiles[i]
		}
	}

	return nil
}

func (h *HoeInfo) GetCurrentScrapingProfile() *HoeProfile {
	return &h.Profiles[0]
}

func (h *HoeInfo) Print() {
	// fmt.Printf("=========== Hoe %s ===========\n", hoe.Profiles[0])
	// fmt.Printf("Url: %s\n", hoe.Url)

	// // Print all fields in Main info
	// fmt.Printf("Name: %s\n", hoe.Name)
	// fmt.Printf("Image url: %s\n", hoe.ImageUrl)
	// fmt.Printf("Price: %s\n", hoe.Price)
	// fmt.Printf("Phone: %s\n", hoe.Phone)
	// fmt.Printf("Address: %s\n", hoe.Address)
	// fmt.Printf("Provider: %s\n", hoe.Provider)
	// fmt.Printf("Status: %s\n", hoe.Status)
	// fmt.Printf("Birth year: %s\n", hoe.BirthYear)
	// fmt.Printf("Height: %s\n", hoe.Height)
	// fmt.Printf("Weight: %s\n", hoe.Weight)
	// fmt.Printf("From: %s\n", hoe.Country)
	// fmt.Printf("Service: %s\n", hoe.Service)
	// fmt.Printf("Work Time: %s\n", hoe.WorkTime)
	// fmt.Printf("Duration: %s\n", hoe.Duration)
	// fmt.Printf("Report count: %d\n", len(hoe.Reports))

	// fmt.Println("==============================")
	fmt.Print("\n\n")
}
