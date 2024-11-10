package models

import "fmt"

type HoeInfo struct {
	ID         string   `json:"id"`
	Url        string   `json:"url"`
	Name       string   `json:"name"`
	ImageUrl   string   `json:"image_url"`
	Price      string   `json:"price"`
	Phone      string   `json:"phone"`
	Address    string   `json:"address"`
	Provider   string   `json:"provider"`
	Status     string   `json:"status"`
	BirthYear  string   `json:"birth_year"`
	Height     string   `json:"height"`
	Weight     string   `json:"weight"`
	From       string   `json:"from"`
	Service    string   `json:"service"`
	Duration   string   `json:"duration"`
	WorkTime   string   `json:"work_time"`
	ReportURLs []string `json:"report_urls"`
}

type HoeReportInfo struct {
	ID          string    `json:"id"`
	DetailUrl   string    `json:"detail_url"`
	Rating      string    `json:"rating"`
	Author      *Reporter `json:"author"`
	Time        string    `json:"time"`
	Description string    `json:"description"`
}

func (hoe *HoeInfo) Print() {
	fmt.Printf("=========== Hoe %s ===========\n", hoe.ID)
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
	fmt.Printf("From: %s\n", hoe.From)
	fmt.Printf("Service: %s\n", hoe.Service)
	fmt.Printf("Duration: %s\n", hoe.Duration)

	fmt.Println("==============================")
	fmt.Print("\n\n")
}
