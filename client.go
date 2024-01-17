package tronscan

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewHttpClient(endpoint string) (c *resty.Client) {
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
		"accept":     "application/json",
	}
	c = resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(time.Second).
		SetBaseURL(endpoint).
		SetHeaders(headers)

	return
}
