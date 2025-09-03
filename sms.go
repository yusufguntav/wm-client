package wmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/yusufguntav/wm-client/models"
)

func (c *Client) SendSmsPreview(req models.PreviewSmsRequest) (models.BulkSmsPreviewResponse, error) {
	buf := &bytes.Buffer{}
	multipartForm := multipart.NewWriter(buf)
	multipartForm.WriteField("numbers", strings.Join(req.Numbers, ","))
	multipartForm.WriteField("message", req.Message)
	multipartForm.WriteField("sender_name", strconv.Itoa(req.SenderName))
	multipartForm.WriteField("send_date", req.SendDate.Format(time.RFC3339))
	multipartForm.WriteField("character_type", strconv.Itoa(int(req.CharacterType)))
	multipartForm.WriteField("message_type", strconv.Itoa(int(req.MessageType)))
	multipartForm.WriteField("add_cancel_link", strconv.FormatBool(req.AddCancelLink))
	err := multipartForm.Close()
	if err != nil {
		return models.BulkSmsPreviewResponse{}, fmt.Errorf("close multipart form error: %w", err)
	}

	respBody, err := c.doFormRequest("POST", "/bulk/preview/sms", multipartForm, buf)
	if err != nil {
		return models.BulkSmsPreviewResponse{}, fmt.Errorf("send sms preview error: %w, response: %s", err, respBody)
	}

	var previewResp models.APIResponse[models.BulkSmsPreviewResponse]
	if err := json.Unmarshal(respBody, &previewResp); err != nil {
		return models.BulkSmsPreviewResponse{}, fmt.Errorf("unmarshal error: %w", err)
	}

	return previewResp.Data, nil
}

func (c *Client) SendSms(req models.SendSmsRequest) (models.SmsSendResponse, error) {
	respBody, err := c.doRequest("POST", "/bulk/sms", req)
	if err != nil {
		return models.SmsSendResponse(""), err
	}

	var resp models.APIResponse[models.SmsSendResponse]
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return models.SmsSendResponse(""), fmt.Errorf("unmarshal error: %w", err)
	}

	return resp.Data, nil
}

func (c *Client) ForceSendSms(req models.PreviewSmsRequest) (models.SmsSendResponse, error) {
	resp, err := c.SendSmsPreview(req)
	if err != nil {
		return models.SmsSendResponse(""), fmt.Errorf("force send sms error: %w", err)
	}

	respSendSms, err := c.SendSms(models.SendSmsRequest{
		ID: resp.ID,
	})
	if err != nil {
		return models.SmsSendResponse(""), fmt.Errorf("force send sms error: %w", err)
	}

	return respSendSms, nil

}
