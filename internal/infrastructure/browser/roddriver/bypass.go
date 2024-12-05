package roddriver

import (
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/cloudflare"
	"github.com/go-rod/rod/lib/proto"
)

func (rb *rodBrowser) BypassCloudflare(url string) (*definitions.BypassResult, error) {
	cloudflareBypasser := cloudflare.NewBypasser("yoori")

	rawResponse, err := cloudflareBypasser.RequestToBypasser(url)
	if err != nil {
		return nil, fmt.Errorf("failed to request to bypasser: %v", err)
	}

	result, err := cloudflareBypasser.ParseResponse(rawResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

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

		err = rb.browser.SetCookies(uniqueCookies)
		if err != nil {
			return nil, fmt.Errorf("failed to set cookies: %v", err)
		}
	}

	return result, nil
}
