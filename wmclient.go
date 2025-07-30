package wmclient

import (
	"net/http"
	"time"
)

const defaultBaseURL = "https://api.toplusms.app"

type Client struct {
	BaseURL    string
	Token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {

	return &Client{
		BaseURL:    defaultBaseURL,
		Token:      token,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}
