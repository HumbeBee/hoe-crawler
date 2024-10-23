package gaito

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/haovoanh28/gai-webscraper/models"
	"github.com/haovoanh28/gai-webscraper/utils"
)

func ProcessDetailPage(ctx context.Context, detailUrl string) (models.Hoe, error) {
	url := BaseUrl + detailUrl

	id, err := strconv.Atoi(utils.GetHoeIDFromUrl(detailUrl))
	if err != nil {
		id = -1
	}

	var mainInfo models.HoeMainInfo
	var detailInfo models.HoeDetailInfo

	err = chromedp.Run(ctx, chromedp.Navigate(url),
		chromedp.WaitVisible(".tab-content"),
		// Get main information
		chromedp.Evaluate(`document.querySelector('.jumbotron .fa.fa-money + span')?.textContent?.replace(/\s+/g, '')`, &mainInfo.Price),
		chromedp.Evaluate(`document.querySelector('.jumbotron .fa.fa-phone + a')?.textContent?.split(".")?.join("").trim()`, &mainInfo.Phone),
		chromedp.Evaluate(`document.querySelector('.jumbotron .fa.fa-map-marker + a')?.textContent?.trim()`, &mainInfo.Address),
		chromedp.Evaluate(`document.querySelector('.jumbotron .fa.fa-user + a')?.textContent?.trim()`, &mainInfo.Author),
		chromedp.Evaluate(`document.querySelector('.jumbotron .fa.fa-file-o + span')?.textContent?.trim()`, &mainInfo.Status),

		// Get detail information
		chromedp.WaitVisible("product-attribute .table-responsive"),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`document.querySelector('product-attribute table > tbody > tr:nth-child(3) > td:nth-child(2) > attribute-dob-box .ng-scope')?.textContent?.trim()`, &detailInfo.BirthYear),
		chromedp.Evaluate(`document.querySelector('product-attribute table > tbody > tr:nth-child(4) > td:nth-child(2) > attribute-number-box .ng-scope')?.textContent?.trim()`, &detailInfo.Height),
		chromedp.Evaluate(`document.querySelector('product-attribute table > tbody > tr:nth-child(5) > td:nth-child(2) > attribute-number-box .ng-scope')?.textContent?.trim()`, &detailInfo.Weight),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('product-attribute table > tbody > tr:nth-child(9) > td:nth-child(2) > attribute-radio-box span span[ng-repeat="item in attributeDto.settings.values"]')).map(item => { const text = item.querySelector('.ng-scope')?.textContent; if (text) { return text.trim(); } return ''; }).filter(Boolean).join(",")`, &detailInfo.From),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('product-attribute table > tbody > tr:nth-child(12) > td:nth-child(2) > attribute-choices-box span span[ng-repeat="item in attributeDto.settings.values"]')).map(item => { const text = item.querySelector('.ng-scope')?.textContent; if (text) { return text.trim(); } return ''; }).filter(Boolean).join(',')`, &detailInfo.Service),
		chromedp.Evaluate(`document.querySelector('product-attribute table > tbody > tr:nth-child(15) > td:nth-child(2) > attribute-text-box span').textContent.trim()`, &detailInfo.Duration),
		chromedp.Evaluate(`document.querySelector('product-attribute table > tbody > tr:nth-child(16) > td:nth-child(2) > attribute-text-box span').textContent.trim()`, &detailInfo.WorkTime),
	)

	if err != nil {
		return models.Hoe{}, err
	}

	fmt.Printf("mainInfo: %+v \n", mainInfo)
	fmt.Printf("detail info: %+v \n", detailInfo)

	return models.Hoe{ID: id, Url: url, MainInfo: &mainInfo, DetailInfo: &detailInfo}, nil
}
