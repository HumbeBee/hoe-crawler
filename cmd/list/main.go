package main

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/cmd/config"
	"github.com/haovoanh28/gai-webscraper/internal/sites/gaito"
)

func main() {
	config.InitLogger()

	scraper := gaito.NewScraper(config.BaseUrl, config.RequestPerSecond, log.Default())
	urlList, err := scraper.GetList()
	if err != nil {
		scraper.ErrorHandler.Fatal(err.Error())
	}

	if len(urlList) == 0 {
		scraper.ErrorHandler.Info("No items found (Maybe Cloudflare block ?)")
	} else {
		log.Println("Found", len(urlList), "items")

		// rateLimiter := browser.NewRateLimiter(1.0)
		// urlChan := make(chan string, len(urlList))
		// resultChan := make(chan models.HoeInfo, len(urlList))
		// var wg sync.WaitGroup

		// numWorkers := 2
		// for i := 0; i < numWorkers; i++ {
		// 	wg.Add(1)
		// 	go func() {
		// 		defer wg.Done()
		// 		for url := range urlChan {
		// 			// Wait for rate limiter before making request
		// 			rateLimiter.Wait()

		// 		}
		// 	}()
		// }

		// go func() {
		// 	wg.Wait()
		// 	close(resultChan)
		// }()

		// for _, url := range urlList {
		// 	urlChan <- url
		// }
	}
}
