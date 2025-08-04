package wmclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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
			return respBody, fmt.Errorf("unexpected status code after token refresh: %d\nResponse:%s", status, respBody)
		}
	}

	return respBody, nil
}

func (c *Client) executeRequest(method, endpoint string, payload interface{}) ([]byte, int, error) {
	var body io.Reader
	contentType := "application/json"

	switch p := payload.(type) {
	case nil:
	case *bytes.Buffer:
		body = p
	case io.Reader:
		body = p
	default:
		if hasMultipartFile(payload) {
			var b bytes.Buffer
			w := multipart.NewWriter(&b)
			if err := writeMultipart(w, payload); err != nil {
				return nil, 0, fmt.Errorf("multipart encode error: %w", err)
			}
			w.Close()
			body = &b
			contentType = w.FormDataContentType()
		} else {
			data, err := json.Marshal(payload)
			if err != nil {
				return nil, 0, fmt.Errorf("json marshal error: %w", err)
			}
			body = bytes.NewBuffer(data)
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), body)
	if err != nil {
		return nil, 0, fmt.Errorf("request creation error: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
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
		return respBody, resp.StatusCode, fmt.Errorf("unexpected status code: %d\nResponse:%s", resp.StatusCode, respBody)
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

func hasMultipartFile(payload interface{}) bool {
	val := reflect.ValueOf(payload)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < val.NumField(); i++ {
		ft := val.Field(i)
		if ft.Kind() == reflect.String && ft.String() != "" {
			// local file path
			tag := val.Type().Field(i).Tag.Get("form")
			if strings.Contains(tag, "file") {
				return true
			}
		}
		if fh, ok := ft.Interface().(*multipart.FileHeader); ok && fh != nil {
			return true
		}
	}
	return false
}

func writeMultipart(w *multipart.Writer, payload interface{}) error {
	val := reflect.ValueOf(payload)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		formTag := fieldType.Tag.Get("form")
		if formTag == "" {
			continue
		}

		switch f := field.Interface().(type) {
		case string:
			// Eğer dosya alanı ve path verilmişse dosyayı oku
			if (strings.Contains(strings.ToLower(formTag), "file") || formTag == "bulk_wp_file") && f != "" {
				file, err := os.Open(f)
				if err != nil {
					return fmt.Errorf("file open error for %s: %w", f, err)
				}
				defer file.Close()

				part, err := w.CreateFormFile(formTag, filepath.Base(f))
				if err != nil {
					return err
				}
				if _, err := io.Copy(part, file); err != nil {
					return err
				}
			} else {
				w.WriteField(formTag, f)
			}
		case bool:
			w.WriteField(formTag, strconv.FormatBool(f))
		case int, int64:
			w.WriteField(formTag, fmt.Sprintf("%v", f))
		case *multipart.FileHeader:
			if f != nil {
				file, err := f.Open()
				if err != nil {
					return err
				}
				defer file.Close()
				part, err := w.CreateFormFile(formTag, f.Filename)
				if err != nil {
					return err
				}
				if _, err := io.Copy(part, file); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
