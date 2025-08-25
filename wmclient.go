package wmclient

import (
	"net/http"
	"time"

	"log"

	"github.com/yusufguntav/wm-client/models"
)

const defaultBaseURL = "https://api.toplusms.app"

type Client struct {
	BaseURL    string
	Token      string
	LoginInfo  models.LoginVerifyCodePayload
	httpClient *http.Client
}

type NewClientArgs struct {
	LoginInfo models.LoginVerifyCodePayload
	BaseURL   string
	AutoLogin bool
}

func NewClient(args NewClientArgs) *Client {
	baseURL := args.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	client := &Client{
		BaseURL:    baseURL,
		LoginInfo:  args.LoginInfo,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

	if args.AutoLogin {
		err := client.RefreshToken()
		if err != nil {
			log.Printf("Failed to refresh token: %v", err)
		}
	}

	return client

}
