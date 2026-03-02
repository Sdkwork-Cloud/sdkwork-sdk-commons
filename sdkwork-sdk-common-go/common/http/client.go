package common

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BaseHttpClient struct {
	config      SdkConfig
	baseURL     string
	timeout     time.Duration
	headers     HttpHeaders
	client      *http.Client
	apiKey      string
	authToken   string
	accessToken string
}

func NewBaseHttpClient(config SdkConfig) *BaseHttpClient {
	timeout := DefaultTimeout
	if config.Timeout != nil && *config.Timeout > 0 {
		timeout = *config.Timeout
	}

	headers := HttpHeaders{}
	for key, value := range config.Headers {
		headers[key] = value
	}

	client := &BaseHttpClient{
		config:      config,
		baseURL:     strings.TrimRight(config.BaseURL, "/"),
		timeout:     time.Duration(timeout) * time.Millisecond,
		headers:     headers,
		client:      &http.Client{Timeout: time.Duration(timeout) * time.Millisecond},
		apiKey:      config.ApiKey,
		authToken:   config.AuthToken,
		accessToken: config.AccessToken,
	}
	client.updateAuthHeaders()
	return client
}

func (c *BaseHttpClient) SetApiKey(apiKey string) {
	c.apiKey = apiKey
	c.updateAuthHeaders()
}

func (c *BaseHttpClient) SetAuthToken(token string) {
	c.authToken = token
	c.updateAuthHeaders()
}

func (c *BaseHttpClient) SetAccessToken(token string) {
	c.accessToken = token
	c.updateAuthHeaders()
}

func (c *BaseHttpClient) SetHeader(key, value string) {
	if c.headers == nil {
		c.headers = HttpHeaders{}
	}
	c.headers[key] = value
}

func (c *BaseHttpClient) Request(method, path string, params map[string]interface{}, body interface{}) ([]byte, error) {
	return c.RequestWithContext(context.Background(), method, path, params, body)
}

func (c *BaseHttpClient) RequestWithContext(ctx context.Context, method, path string, params map[string]interface{}, body interface{}) ([]byte, error) {
	requestURL, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, NewNetworkError(err.Error())
	}

	if len(params) > 0 {
		query := requestURL.Query()
		for key, value := range params {
			if value != nil {
				query.Set(key, stringify(value))
			}
		}
		requestURL.RawQuery = query.Encode()
	}

	var requestBody io.Reader
	if body != nil {
		data, marshalErr := json.Marshal(body)
		if marshalErr != nil {
			return nil, NewValidationError(marshalErr.Error(), nil)
		}
		requestBody = bytes.NewBuffer(data)
	}

	req, reqErr := http.NewRequestWithContext(ctx, method, requestURL.String(), requestBody)
	if reqErr != nil {
		return nil, NewNetworkError(reqErr.Error())
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, doErr := c.client.Do(req)
	if doErr != nil {
		if ctx.Err() == context.DeadlineExceeded || doErr == context.DeadlineExceeded {
			timeout := int(c.timeout.Milliseconds())
			return nil, NewTimeoutError(doErr.Error(), &timeout)
		}
		return nil, NewNetworkError(doErr.Error())
	}
	defer resp.Body.Close()

	responseBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, NewNetworkError(readErr.Error())
	}

	if resp.StatusCode >= 400 {
		return nil, c.handleHttpError(resp.StatusCode, responseBody, resp.Header)
	}

	return responseBody, nil
}

func (c *BaseHttpClient) Get(path string, params map[string]interface{}) ([]byte, error) {
	return c.Request(http.MethodGet, path, params, nil)
}

func (c *BaseHttpClient) Post(path string, body interface{}) ([]byte, error) {
	return c.Request(http.MethodPost, path, nil, body)
}

func (c *BaseHttpClient) Put(path string, body interface{}) ([]byte, error) {
	return c.Request(http.MethodPut, path, nil, body)
}

func (c *BaseHttpClient) Delete(path string, params map[string]interface{}) ([]byte, error) {
	return c.Request(http.MethodDelete, path, params, nil)
}

func (c *BaseHttpClient) Patch(path string, body interface{}) ([]byte, error) {
	return c.Request(http.MethodPatch, path, nil, body)
}

func (c *BaseHttpClient) updateAuthHeaders() {
	delete(c.headers, "Authorization")
	delete(c.headers, "Access-Token")
	delete(c.headers, "X-API-Key")

	if c.apiKey != "" {
		c.headers["Authorization"] = "Bearer " + c.apiKey
	}
	if c.authToken != "" {
		c.headers["Authorization"] = "Bearer " + c.authToken
	}
	if c.accessToken != "" {
		c.headers["Access-Token"] = c.accessToken
	}
}

func (c *BaseHttpClient) handleHttpError(status int, responseBody []byte, headers http.Header) error {
	message := strings.TrimSpace(string(responseBody))
	if message == "" {
		message = "HTTP error"
	}

	switch status {
	case http.StatusUnauthorized:
		return NewAuthenticationError(message)
	case http.StatusForbidden:
		return NewForbiddenError(message)
	case http.StatusNotFound:
		return NewNotFoundError(message)
	case http.StatusBadRequest:
		return NewValidationError(message, nil)
	case http.StatusTooManyRequests:
		var retryAfter *int
		if value := headers.Get("Retry-After"); value != "" {
			if parsed, err := strconv.Atoi(value); err == nil {
				retryAfter = &parsed
			}
		}
		return NewRateLimitError(message, retryAfter)
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return NewServerError(message, status)
	default:
		return &SdkError{
			Message:    message,
			Code:       ErrCodeUnknown,
			HttpStatus: &status,
			Timestamp:  Now(),
		}
	}
}

func stringify(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return strings.TrimSpace(strings.ReplaceAll(strings.Trim(fmtJSON(v), `"`), "\n", ""))
	}
}

func fmtJSON(value interface{}) string {
	data, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(data)
}
