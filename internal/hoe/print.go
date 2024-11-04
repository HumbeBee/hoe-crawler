package hoe

import "fmt"

func (hoe *Hoe) Print() {
	fmt.Printf("=========== Hoe %s ===========\n", hoe.ID)
	fmt.Printf("Url: %s\n", hoe.Url)

	// Print all fields in Main info
	fmt.Printf("\nMain info:\n")
	fmt.Printf("Name: %s\n", hoe.MainInfo.Name)
	fmt.Printf("Image url: %s\n", hoe.MainInfo.ImageUrl)
	fmt.Printf("Price: %s\n", hoe.MainInfo.Price)
	fmt.Printf("Phone: %s\n", hoe.MainInfo.Phone)
	fmt.Printf("Address: %s\n", hoe.MainInfo.Address)
	fmt.Printf("Author: %s\n", hoe.MainInfo.Author)
	fmt.Printf("Status: %s\n", hoe.MainInfo.Status)

	// Print all fields in Detail info
	fmt.Printf("\nDetail info:\n")
	fmt.Printf("Birth year: %s\n", hoe.DetailInfo.BirthYear)
	fmt.Printf("Height: %s\n", hoe.DetailInfo.Height)
	fmt.Printf("Weight: %s\n", hoe.DetailInfo.Weight)
	fmt.Printf("From: %s\n", hoe.DetailInfo.From)
	fmt.Printf("Service: %s\n", hoe.DetailInfo.Service)
	fmt.Printf("Duration: %s\n", hoe.DetailInfo.Duration)

	// Print all fields in Report info
	fmt.Printf("\nReport info:\n")
	for _, report := range hoe.DetailReports {
		fmt.Printf("Report Url: %s\n", report.DetailUrl)
		fmt.Printf("Report Rating: %s\n", report.Rating)
		fmt.Printf("Report Author: %s\n", report.Author)
		fmt.Printf("Report Description: %s\n", report.Description)
	}

	fmt.Println("==============================")
	fmt.Print("\n\n")
}
