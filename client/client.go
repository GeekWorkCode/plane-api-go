package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/GeekWorkCode/plane-api-go/models"
)

const (
	defaultBaseURL = "https://app.plane.so/api/v1"
	userAgent      = "plane-api-go/0.1.0"
)

// Client represents the Plane API client
type Client struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
	userAgent  string
	debug      bool
}

// NewClient creates a new Plane API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		baseURL:    "https://api.plane.so/api/v1",
		userAgent:  userAgent,
		debug:      false,
	}
}

// SetDebug enables or disables debug logging
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// SetBaseURL sets the base URL for API requests
func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = strings.TrimRight(baseURL, "/")
}

// NewRequest creates a new API request
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	url := c.baseURL + path

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-API-Key", c.apiKey)

	return req, nil
}

// Do sends an API request and returns the API response
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			fmt.Printf("REQUEST:\n%s\n", string(dump))
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			fmt.Printf("RESPONSE:\n%s\n", string(dump))
		}
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp, fmt.Errorf("API error: %s (Status: %d)", req.URL.String(), resp.StatusCode)
		}

		if c.debug {
			fmt.Printf("ERROR RESPONSE BODY: %s\n", string(bodyBytes))
		}

		var errorResp models.ErrorResponse
		if err := json.Unmarshal(bodyBytes, &errorResp); err != nil {
			return resp, fmt.Errorf("API error: %s (Status: %d)\nBody: %s", req.URL.String(), resp.StatusCode, string(bodyBytes))
		}

		if errorResp.Error != "" {
			return resp, fmt.Errorf("API error: %s (Status: %d)\nError: %s", req.URL.String(), resp.StatusCode, errorResp.Error)
		}
		if errorResp.Message != "" {
			return resp, fmt.Errorf("API error: %s (Status: %d)\nMessage: %s", req.URL.String(), resp.StatusCode, errorResp.Message)
		}

		return resp, fmt.Errorf("API error: %s (Status: %d)\nBody: %s", req.URL.String(), resp.StatusCode, string(bodyBytes))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr != nil && decErr != io.EOF {
				return resp, decErr
			}
		}
	}

	return resp, nil
}
