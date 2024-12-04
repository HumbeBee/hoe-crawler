package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/HumbeBee/hoe-crawler/internal/definitions"
	"io"
	"net/http"
	"time"

	"github.com/HumbeBee/hoe-crawler/internal/infrastructure/interfaces"
)

type yooriBypasser struct {
	baseUrl  string
	basePort int
	client   *http.Client
}

type YooriResponse struct {
	Status   string  `json:"status"`
	Message  string  `json:"message"`
	StartTS  float64 `json:"startTimestamp"`
	EndTS    float64 `json:"endTimestamp"`
	Solution struct {
		Status  string `json:"status"`
		URL     string `json:"url"`
		Cookies []struct {
			Name   string `json:"name"`
			Value  string `json:"value"`
			Domain string `json:"domain"`
			Path   string `json:"path"`
			Secure bool   `json:"secure"`
		} `json:"cookies"`
		UserAgent string  `json:"userAgent"`
		Response  *string `json:"response"`
	} `json:"solution"`
}

func newYooriBypasser() interfaces.CloudflareBypasser {
	return &yooriBypasser{
		baseUrl:  "http://localhost",
		basePort: 20080,
		client: &http.Client{
			Timeout: 3 * time.Minute,
		},
	}
}

func (y *yooriBypasser) RequestToBypasser(url string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s:%d/v1", y.baseUrl, y.basePort)

	data := map[string]interface{}{
		"cmd":        "request.get",
		"url":        url,
		"maxTimeout": 60000,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := y.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func (y *yooriBypasser) ParseResponse(rawResponse []byte) (*definitions.BypassResult, error) {
	var yooriResp YooriResponse
	if err := json.Unmarshal(rawResponse, &yooriResp); err != nil {
		return nil, fmt.Errorf("failed to parse Yoori response: %v", err)
	}

	var cookies []*definitions.Cookie
	for _, cookie := range yooriResp.Solution.Cookies {
		cookies = append(cookies, &definitions.Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Domain: cookie.Domain,
			Path:   cookie.Path,
			Secure: cookie.Secure,
		})
	}

	return &definitions.BypassResult{
		Success:             yooriResp.Status == "ok",
		IsChallengeDetected: yooriResp.Message != "Challenge not detected!",
		Cookies:             cookies,
		UserAgent:           yooriResp.Solution.UserAgent,
	}, nil
}
