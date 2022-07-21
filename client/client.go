package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	mimeType         = "application/json"
	userAgent        = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	defaultRateLimit = 4 // 4 rps
)

const (
	DefaultURL = "https://api.fortnox.se/3/"
	TestURL    = "https://api.fortnox.se/3/test/"
)

const (
	defaultTimeout = 10 * time.Second
)

var (
	defaultHeaders = map[string]string{
		"Accept":       mimeType,
		"Content-Type": mimeType,
		"User-Agent":   userAgent,
	}
)

type Client struct {
	clientOptions *Options
}

func NewClient(options ...OptionFunc) *Client {
	c := &http.Client{
		Timeout: defaultTimeout,
	}

	co := &Options{
		BaseURL:    DefaultURL,
		HTTPClient: c,
	}

	for _, f := range options {
		f(co)
	}

	return &Client{
		clientOptions: co,
	}
}

func (c *Client) buildURL(section string) (*url.URL, error) {
	u, err := url.Parse(c.clientOptions.BaseURL)
	if err != nil {
		return nil, err
	}

	u2, err := url.Parse(section)
	if err != nil {
		return nil, err
	}

	return u.ResolveReference(u2), nil
}

func (c *Client) _GET(ctx context.Context, uri string, params url.Values, resp interface{}) error {
	return c.request(ctx, http.MethodGet, uri, params, nil, resp)
}

func (c *Client) _POST(ctx context.Context, uri string, params url.Values, body, resp interface{}) error {
	return c.request(ctx, http.MethodPost, uri, params, body, resp)
}

func (c *Client) _PUT(ctx context.Context, uri string, params url.Values, body, resp interface{}) error {
	return c.request(ctx, http.MethodPut, uri, params, body, resp)
}

func (c *Client) _DELETE(ctx context.Context, uri string) error {
	return c.request(ctx, http.MethodDelete, uri, nil, nil, nil)
}

func (c *Client) _DELETEWithResult(ctx context.Context, uri string, resp interface{}) error {
	return c.request(ctx, http.MethodDelete, uri, nil, nil, resp)
}

func (c *Client) request(
	ctx context.Context,
	method string,
	uri string,
	params url.Values,
	body, result interface{}) error {
	u, err := c.buildURL(uri)
	if err != nil {
		return err
	}

	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.clientOptions.AccessToken),
		"Client-Secret": c.clientOptions.ClientSecret,
	}

	if strings.ToLower(method) == "delete" {
		bodyBuffer := http.NoBody
		return request(ctx, c.clientOptions.HTTPClient, headers, method, u.String(), bodyBuffer, result)
	}

	bodyBuffer := &bytes.Buffer{}
	err = json.NewEncoder(bodyBuffer).Encode(body)
	if err != nil {
		return err
	}

	return request(ctx, c.clientOptions.HTTPClient, headers, method, u.String(), bodyBuffer, result)
}

func request(
	ctx context.Context,
	client *http.Client,
	headers map[string]string,
	method string,
	url string,
	data io.Reader,
	result interface{}) error {

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return fmt.Errorf("%s: %s", ErrCreateRequest, err)
	}

	req = req.WithContext(ctx)

	for k, v := range defaultHeaders {
		req.Header.Add(k, v)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("%s: %s", ErrSendRequest, err)
	}

	defer func() {
		_, _ = io.CopyN(ioutil.Discard, resp.Body, 64)
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case 200, 201:
		bodyPreview, _ := getRespBodyPreview(resp, 30)
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to decode json from response [%s]", bodyPreview))
		}
		return nil
	case 204:
		return nil
	default:
		// if malformed, want to see
		errMsg := &ErrorResp{}
		bodyPreview, _ := getRespBodyPreview(resp, 128)
		if err := json.NewDecoder(resp.Body).Decode(&errMsg); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to decode %d error from response [%s]", resp.StatusCode, bodyPreview))
		}
		return FortnoxError{
			HTTPStatus: resp.StatusCode,
			Code:       errMsg.ErrorInformation.Code,
			Message:    errMsg.ErrorInformation.Message,
		}
	}

}

// Get a preview of the body (without affecting the resp.Body reader)
func getRespBodyPreview(resp *http.Response, len int64) (string, error) {

	part, err := ioutil.ReadAll(io.LimitReader(resp.Body, len))
	if err != nil {
		return "", err
	}

	// recombine the buffered part of the body with the rest of the stream
	resp.Body = ioutil.NopCloser(io.MultiReader(bytes.NewReader(part), resp.Body))
	return string(part), nil
}

var (
	ErrCreateRequest = errors.New("error creating request")
	ErrSendRequest   = errors.New("error sending request")
)

// ErrorResp error response from Fortnox
type ErrorResp struct {
	ErrorInformation ErrorMessage
}

// FortnoxError internal Fortnox's error holder
type FortnoxError struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (f FortnoxError) Error() string {
	return fmt.Sprintf("%d - %s", f.Code, f.Message)
}

// Translate translates error message to [langCode]
// TODO: would be awesome to translate error from Sweden to [langCode]
func (f FortnoxError) Translate(langCode string) string {
	return ""
}

// ErrorMessage response type
type ErrorMessage struct {
	Error   int    `json:"Error"`
	Message string `json:"Message"`
	Code    int    `json:"Code"`
}
