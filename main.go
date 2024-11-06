package main

import (
	"fmt"
	"sync"

	"github.com/haovoanh28/gai-webscraper/internal/models"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	urlList := gaito.ProcessListPage()

	rateLimiter := browser.NewRateLimiter(1.0)
	urlChan := make(chan string, len(urlList))
	resultChan := make(chan models.HoeInfo, len(urlList))
	var wg sync.WaitGroup

	numWorkers := 2
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlChan {
				// Wait for rate limiter before making request
				rateLimiter.Wait()

				if result := gaito.ProcessDetailUrl(url); result != nil {
					fmt.Println("Processed", url)
					resultChan <- *result
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Test for first 4 item
	for _, url := range urlList {
		urlChan <- url
	}
	close(urlChan)

	var hoeList []models.HoeInfo
	for hoe := range resultChan {
		hoeList = append(hoeList, hoe)
	}

	for _, hoe := range hoeList {
		hoe.Print()

		if len(hoe.ReportURLs) > 0 {
			// Process report url or put it into db ???
		}
	}

}
