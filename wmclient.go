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

func NewClient(loginInfo models.LoginVerifyCodePayload) *Client {

	client := &Client{
		BaseURL:    defaultBaseURL,
		LoginInfo:  loginInfo,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

	err := client.RefreshToken()
	if err != nil {
		log.Printf("Failed to refresh token: %v", err)
	}

	return client

}
