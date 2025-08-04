package wmclient

import (
	"encoding/json"
	"fmt"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) LoginVerifyCode(payload models.LoginVerifyCodePayload) (models.VerificationData, error) {
	respBody, err := c.doRequest("POST", "/login/sms", payload)
	if err != nil {
		return models.VerificationData{}, fmt.Errorf("login verify code error: %w, response: %s", err, respBody)
	}

	var verificationResp models.APIResponse[[]models.VerificationData]
	if err := json.Unmarshal(respBody, &verificationResp); err != nil {
		return models.VerificationData{}, fmt.Errorf("unmarshal error: %w", err)
	}

	if len(verificationResp.Data) == 0 {
		return models.VerificationData{}, fmt.Errorf("no verification data received")
	}

	return verificationResp.Data[0], nil
}

func (c *Client) Login(payload models.LoginPayload) (models.LoginResponse, error) {
	respBody, err := c.doRequest("POST", "/login", payload)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("login error: %w, response: %s", err, respBody)
	}

	var loginResp models.APIResponse[models.LoginResponse]
	if err := json.Unmarshal(respBody, &loginResp); err != nil {
		return models.LoginResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return loginResp.Data, nil
}
