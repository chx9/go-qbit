package qbit

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"

	wrapper "github.com/pkg/errors"
)

type Client struct {
	baseURL    string
	HTTPClient *http.Client
	Authed     bool
}

func NewClient(baseURL string) *Client {
	client := &Client{}

	if baseURL[len(baseURL)-1:] == "/" {
		baseURL += baseURL[0 : len(baseURL)-1]
	}

	client.baseURL = baseURL

	jar, _ := cookiejar.New(nil)
	client.HTTPClient = &http.Client{
		Jar: jar,
	}
	return client
}

func (client *Client) Login(username string, password string) error {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequest(
		"POST",
		client.baseURL+"/api/v2/auth/login",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", client.baseURL)
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("login failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed, status code: %d", resp.StatusCode)
	}
	cookies := resp.Cookies()
	jar, _ := cookiejar.New(nil)
	parsedURL, _ := url.Parse(client.baseURL)
	for _, cookie := range cookies {
		jar.SetCookies(parsedURL, []*http.Cookie{cookie})
	}
	client.HTTPClient = &http.Client{
		Jar: jar,
	}
	client.Authed = true
	return nil
}

func (client *Client) Logout() error {
	_, err := client.Get("logout", nil)
	if err != nil {
		return err
	}
	client.Authed = false
	return nil
}

func (client *Client) Get(endpoint string, opts map[string]string) (*http.Response, error) {
	if !client.Authed {
		return nil, NotLogin
	}
	req, err := http.NewRequest("GET", client.baseURL+endpoint, nil)
	if err != nil {
		return nil, FailedToBuildRequest(err)
	}
	req.Header.Set("User-Agent", "qbit-client")

	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, RequestFailed(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, wrapper.Errorf("API request %s, failed: status %d", endpoint, resp.StatusCode)
	}
	return resp, nil
}

func (client *Client) PostURL(endpoint string, opts map[string]string) (*http.Response, error) {
	if !client.Authed {
		return nil, NotLogin
	}
	req, err := http.NewRequest("Post", client.baseURL+endpoint, nil)
	if err != nil {
		return nil, FailedToBuildRequest(err)
	}
	req.Header.Set("User-Agent", "qbit-client")

	if opts != nil {
		query := req.URL.Query()
		for k, v := range opts {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, RequestFailed(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, wrapper.Errorf("API request %s, failed: status %d", endpoint, resp.StatusCode)
	}
	return resp, nil
}

func (client *Client) Post(endpoint string, opts map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", client.baseURL+endpoint, nil)
	if err != nil {
		return nil, FailedToBuildRequest(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "qbit-client")
	if opts != nil {
		form := url.Values{}
		for k, v := range opts {
			form.Add(k, v)
		}
		req.PostForm = form
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, RequestFailed(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, wrapper.Errorf("API request %s, failed: status %d", endpoint, resp.StatusCode)
	}
	return resp, nil
}

func (client *Client) PostFormData(endpoint string, formData map[string]string) (*http.Response, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	for key, value := range formData {
		if err := writer.WriteField(key, value); err != nil {
			return nil, wrapper.Wrap(err, "failed to write form field")
		}
	}

	if err := writer.Close(); err != nil {
		return nil, wrapper.Wrap(err, "failed to close multipart writer")
	}

	req, err := http.NewRequest("POST", client.baseURL+endpoint, &buffer)
	if err != nil {
		return nil, FailedToBuildRequest(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "qbit-client")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, RequestFailed(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, wrapper.Errorf("API request %s, failed: status %d", endpoint, resp.StatusCode)
	}
	return resp, nil
}

func (client *Client) PostFileWithForm(endpoint string, formData map[string]string, fileFieldName, filePath string) (*http.Response, error) {
	if !client.Authed {
		return nil, NotLogin
	}
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, wrapper.Wrap(err, "error opening file")
	}
	defer file.Close()

	formWriter, err := writer.CreateFormFile(fileFieldName, path.Base(filePath))
	if err != nil {
		return nil, wrapper.Wrap(err, "error creating file form field")
	}

	if _, err = io.Copy(formWriter, file); err != nil {
		return nil, wrapper.Wrap(err, "error copying file contents")
	}

	for key, value := range formData {
		if err := writer.WriteField(key, value); err != nil {
			return nil, wrapper.Wrap(err, "failed to write form field")
		}
	}

	if err := writer.Close(); err != nil {
		return nil, wrapper.Wrap(err, "failed to close multipart writer")
	}

	req, err := http.NewRequest("POST", client.baseURL+endpoint, &buffer)
	if err != nil {
		return nil, FailedToBuildRequest(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", "qbit-client")

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, RequestFailed(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, wrapper.Errorf("API request %s, failed: status %d", endpoint, resp.StatusCode)
	}
	return resp, nil
}
