package roddriver

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/cloudflare"
	"github.com/HumbeBee/hoe-crawler/internal/utils"
	"github.com/go-rod/rod/lib/proto"
)

func (rb *rodBrowser) BypassCloudflare(url string) (*definitions.BypassResult, error) {
	cloudflareBypasser := cloudflare.NewBypasser("yoori")

	rawResponse, err := cloudflareBypasser.RequestToBypasser(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request to bypasser: %v", err)
	}

	fmt.Println("Raw response:")
	fmt.Println(string(rawResponse))

	result, err := cloudflareBypasser.ParseResponse(rawResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	fmt.Println("Parsed response:")
	utils.PrintJSON(result)
	if result.Success {
		cookieMap := make(map[string]*proto.NetworkCookieParam)
		for _, cookie := range result.Cookies {
			cookieMap[cookie.Name] = &proto.NetworkCookieParam{
				Name:   cookie.Name,
				Value:  cookie.Value,
				Domain: cookie.Domain,
				Path:   cookie.Path,
				Secure: cookie.Secure,
			}
		}

		var uniqueCookies []*proto.NetworkCookieParam
		for _, cookie := range cookieMap {
			uniqueCookies = append(uniqueCookies, cookie)
		}

		for _, cookie := range uniqueCookies {
			fmt.Printf("Setting cookie: %+v\n", cookie)
		}

		err = rb.browser.SetCookies(uniqueCookies)
		if err != nil {
			return nil, fmt.Errorf("failed to set cookies: %v", err)
		}
	}

	return result, nil
}
