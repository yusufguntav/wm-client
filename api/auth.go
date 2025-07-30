package wmclient

import (
	"encoding/json"
	"fmt"
)

// LoginVerifyCode sends login credentials and returns success or error.
func (c *Client) LoginVerifyCode(payload LoginVerifyCodePayload) error {
	respBody, err := c.doRequest("POST", "/login/sms", payload)
	if err != nil {
		return fmt.Errorf("login verify code error: %w, response: %s", err, respBody)
	}

	return nil
}

// Login completes login using verify code and returns the token.
func (c *Client) Login(payload LoginPayload) (string, error) {
	respBody, err := c.doRequest("POST", "/login", payload)
	if err != nil {
		return "", fmt.Errorf("login error: %w, response: %s", err, respBody)
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(respBody, &loginResp); err != nil {
		return "", fmt.Errorf("unmarshal error: %w", err)
	}

	// Token'i client üzerine set et
	c.Token = loginResp.Token

	return loginResp.Token, nil
}
