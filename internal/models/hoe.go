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

func (h *HoeInfo) GetCurrentScrapingProfile() (*HoeProfile, error) {
	if len(h.Profiles) == 0 {
		return nil, fmt.Errorf("hoe has no profile")
	}

	// Why Profiles[0] ?
	// Because the first profile is the one that is currently being scraped
	return &h.Profiles[0], nil
}

func (h *HoeInfo) Print() {
	fmt.Printf("=========== Hoe Info ===========\n")
	fmt.Printf("Name: %s\n", h.Name)
	fmt.Printf("Phone: %s\n", h.Phone)
	fmt.Printf("Birth Year: %s\n", h.BirthYear)
	fmt.Printf("Height: %s\n", h.Height)
	fmt.Printf("Weight: %s\n", h.Weight)
	fmt.Printf("Country: %s\n", h.Country)
	fmt.Printf("Last Scraped At: %s\n", h.LastScrapedAt)

	if len(h.Profiles) > 0 {
		fmt.Println("\n--- Latest Profile ---")
		profile := h.Profiles[0]
		fmt.Printf("Site: %s\n", profile.Site)
		fmt.Printf("Original ID: %d\n", profile.OriginID)
		fmt.Printf("URL: %s\n", profile.Url)
		fmt.Printf("Image URL: %s\n", profile.ImageUrl)
		fmt.Printf("Price: %s\n", profile.Price)
		fmt.Printf("Area: %s\n", profile.Area)
		fmt.Printf("Provider: %s\n", profile.Provider)
		fmt.Printf("Status: %s\n", profile.Status)
		fmt.Printf("Service: %s\n", profile.Service)
		fmt.Printf("Duration: %s\n", profile.Duration)
		fmt.Printf("Work Time: %s\n", profile.WorkTime)
		fmt.Printf("Report Count: %d\n", len(profile.Reports))
	}

	fmt.Printf("\nTotal Profiles: %d\n", len(h.Profiles))
	fmt.Printf("Total Working Histories: %d\n", len(h.WorkingHistories))
	fmt.Println("==============================")
	fmt.Print("\n\n")
}
