package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/khanhct/go-lib-core/sdk/response"
)


type HttpClient struct {
	BaseURL string
	Headers map[string]string
	Client *http.Client
}

func NewClient(baseUrl string, headers map[string]string) *HttpClient{
	return &HttpClient{
		BaseURL: baseUrl,
		Headers: headers,
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *HttpClient) sendRequest(req *http.Request, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	newHeaders := c.Headers
	if headers != nil {
		for key, val := range headers{
			newHeaders[key] = val
		}
	}

	for key, val := range newHeaders {
		fmt.Println(fmt.Sprintf("%s %s", key, val))
		req.Header.Set(key, val)
	}

	if timeout != 0 {
		c.Client.Timeout = timeout
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	resp := response.HttpResponse{}
	resp.Code = res.StatusCode

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(io.Reader(res.Body)).Decode(&resp); err != nil {
		return &resp, err
	}

	return &resp, nil
}

func (c *HttpClient) GET(path string, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, headers, timeout)
}

func (c *HttpClient) POST(path string, data interface{}, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.BaseURL, path), bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, headers, timeout)
}

func (c *HttpClient) PUT(path string, data interface{}, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.BaseURL, path), bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, headers, timeout)
}

func (c *HttpClient) PATCH(path string, data interface{}, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	jsonData, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.BaseURL, path), bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, headers, timeout)
}

func (c *HttpClient) DELETE(path string, headers map[string]string, timeout time.Duration) (*response.HttpResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.BaseURL, path), nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req, headers, timeout)
}
