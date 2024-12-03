package roddriver

import "github.com/HumbeBee/hoe-crawler/internal/infrastructure/cloudflare"

func (b *rodBrowser) BypassCloudflare(url string) error {
	cloudflareBypasser := cloudflare.NewBypasser("yoori")
	cloudflareBypasser.GetCookies(url)

	return nil
}
