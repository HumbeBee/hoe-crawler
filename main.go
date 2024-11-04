package main

import (
	"fmt"
	"sync"

	"github.com/haovoanh28/gai-webscraper/gaito"
	"github.com/haovoanh28/gai-webscraper/internal/hoe"
	"github.com/haovoanh28/gai-webscraper/utils"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	urlList := gaito.ProcessListPage()

	rateLimiter := utils.NewRateLimiter(1.0)
	urlChan := make(chan string, len(urlList))
	resultChan := make(chan hoe.Hoe, len(urlList))
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
	for i := 0; i < 30; i++ {
		urlChan <- urlList[i]
	}
	close(urlChan)

	var results []hoe.Hoe
	for result := range resultChan {
		results = append(results, result)
	}

	for _, result := range results {
		result.Print()
	}

	// if len(hoe.ReportURLs) > 0 {
	// 	urlChans := make(chan string, len(hoe.ReportURLs))

	// 	for _, reportUrl := range hoe.ReportURLs {
	// 		urlChans <- reportUrl
	// 		// reportDetail := gaito.ProcessReportPage(reportUrl)
	// 	}
	// 	close(urlChans)
	// }
}
