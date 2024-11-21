package gaito

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/go-rod/rod"
// 	"github.com/haovoanh28/gai-webscraper/internal/models"
// 	"github.com/haovoanh28/gai-webscraper/internal/utils"
// 	"github.com/haovoanh28/gai-webscraper/internal/utils/browser"
// )

// func ProcessReportPage(reportUrl string) models.HoeReportInfo {
// 	url := BaseUrl + reportUrl
// 	id := utils.GetIDFromUrl(reportUrl)

// 	reportInfo := models.HoeReportInfo{
// 		ID:        id,
// 		DetailUrl: reportUrl,
// 	}

// 	page := rod.New().MustConnect().MustPage(url).MustWaitStable()
// 	element := page.MustElement(`review-detail-cmp`).MustWaitVisible()
// 	defer page.Close()

// 	stars, err := page.Elements(`span[ng-model="data.review.score"] i.fa-heart[ng-repeat="r in range track by $index"]`)
// 	if err != nil {
// 		panic(fmt.Errorf(`failed to get stars: %v`, err))
// 	}
// 	reportInfo.Rating = strconv.Itoa(len(stars))

// 	reportInfo.Description = browser.GetElementText(element, `div[ng-switch-when="textarea"] span[ng-bind="elem.question.response"]`, id+"_report_description")

// 	// Process author
// 	authorSectionElement, err := page.Element(`div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > div.col-md-4 > div:nth-child(2) > div > div.ow_user_list_data`)
// 	if err != nil {
// 		panic(fmt.Errorf(`failed to get author section: %v`, err))
// 	}
// 	authorUrlElement := authorSectionElement.MustElement(`a.ng-binding`)
// 	authorUrl := authorUrlElement.MustAttribute(`href`)
// 	reportInfo.Author = &models.Reporter{
// 		ID:   utils.GetIDFromUrl(*authorUrl),
// 		Name: authorUrlElement.MustText(),
// 		Url:  *authorUrl,
// 	}
// 	reportInfo.Time = browser.GetElementText(element, `div.ow_page_padding > div > div > div > div > div > div:nth-child(3) > div > div.col-md-8 > review-detail-cmp > div:nth-child(1) > em > small`, id+"_report_time")

// 	return reportInfo
// }
