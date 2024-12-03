package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func newYooriBypasser() interfaces.CloudflareBypasser {
	return &yooriBypasser{
		baseUrl:  "http://localhost",
		basePort: 20080,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (y *yooriBypasser) GetCookies(url string) ([]byte, error) {
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

	return io.ReadAll(resp.Body)
}
