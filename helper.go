package wmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) doRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	respBody, status, err := c.executeRequest(method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	if status == http.StatusForbidden {
		// Token yenile
		var err error
		for i := 0; i < 3; i++ {
			err = c.RefreshToken()
			if err == nil {
				break
			}
			time.Sleep(time.Second * 2)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to refresh token after 3 attempts: %w", err)
		}

		// İsteği yeniden yap
		respBody, status, err := c.executeRequest(method, endpoint, payload)

		if err != nil {
			return nil, fmt.Errorf("request after token refresh failed: %w", err)
		}

		if status < 200 || status >= 300 {
			return respBody, fmt.Errorf("unexpected status code after token refresh: %d", status)
		}
	}

	return respBody, nil
}

func (c *Client) executeRequest(method, endpoint string, payload interface{}) ([]byte, int, error) {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, 0, fmt.Errorf("json marshal error: %w", err)
		}
		body = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), body)
	if err != nil {
		return nil, 0, fmt.Errorf("request creation error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request execution error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("read response error: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, resp.StatusCode, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return respBody, resp.StatusCode, nil
}

func (c *Client) RefreshToken() error {
	resp, err := c.LoginVerifyCode(c.LoginInfo)

	if err != nil {
		log.Printf("Failed to get verification code: %v", err)
	}

	loginResp, err := c.Login(models.LoginPayload{
		Identifier: c.LoginInfo.Identifier,
		VerifyCode: resp.VerificationCode,
	})

	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}

	c.Token = loginResp.Token
	return nil
}
