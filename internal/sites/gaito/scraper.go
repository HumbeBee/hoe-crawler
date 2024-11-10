package gaito

import (
	"log"

	"github.com/haovoanh28/gai-webscraper/internal/utils/errutil"
)

type Scraper struct {
	baseURL        string
	requestsPerSec float64
	ErrorHandler   *errutil.ErrorHandler
}

func NewScraper(baseURL string, requestsPerSec float64, logger *log.Logger) *Scraper {
	return &Scraper{baseURL: baseURL, requestsPerSec: requestsPerSec, ErrorHandler: errutil.NewErrorHandler(logger, errutil.DEBUG)}
}

func (s *Scraper) ProcessListPage() {
	urlList, err := s.processListPage()

	if err != nil {
		s.ErrorHandler.Fatal(err.Error())
	}

	if len(urlList) == 0 {
		s.ErrorHandler.Info("No items found (Maybe Cloudflare block ?)")
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

func (s *Scraper) ProcessDetailPage(url string) {
	hoeInfo, err := s.getHoeInfoFromDetailPage(url)
	if err != nil {
		s.ErrorHandler.Fatal(err.Error())
	}

	// Save hoeInfo to db

	if len(hoeInfo.ReportUrls) > 0 {
		// Should save urls into db
	}

	hoeInfo.Print()
}
