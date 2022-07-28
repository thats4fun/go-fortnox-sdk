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
	refreshTokenURL = "https://apps.fortnox.se/oauth-v1/token"
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

type Scope string

const (
	Salary             Scope = "salary"
	Bookkeeping        Scope = "bookkeeping"
	Archive            Scope = "archive"
	FileConnection     Scope = "connectfile"
	Article            Scope = "article"
	CompanyInformation Scope = "companyinformation"
	Settings           Scope = "settings"
	Invoice            Scope = "invoice"
	CostCenters        Scope = "costcenter"
	Currency           Scope = "currency"
	Customer           Scope = "customer"
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

func (c *Client) String() string {
	return fmt.Sprintf("AccessToken: %s, RefreshToken: %s, ClientSecret:%s, BaseURL: %s",
		c.clientOptions.AccessToken, c.clientOptions.RefreshToken, c.clientOptions.ClientSecret, c.clientOptions.BaseURL)
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

func (c *Client) RefreshToken() error {
	body := &bytes.Buffer{}

	body.WriteString(
		fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", c.clientOptions.RefreshToken),
	)

	resp, err := http.Post(refreshTokenURL, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	tokenInfo := TokenInfo{}
	err = json.Unmarshal(bts, &tokenInfo)
	if err != nil {
		return err
	}

	c.clientOptions.AccessToken = tokenInfo.AccessToken
	c.clientOptions.RefreshToken = tokenInfo.RefreshToken

	return nil
}

type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
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
			if err == io.EOF {
				return nil
			}
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
		msg := errMsg.ErrorInformation.Message
		if errMsg.ErrorInformation.Code == 0 {
			msg = fmt.Sprintf("%s | try to refresh token", bodyPreview)
		}
		return FortnoxError{
			HTTPStatus: resp.StatusCode,
			Code:       errMsg.ErrorInformation.Code,
			Message:    msg,
		}
	}
}

type MessageWrapper struct {
	Message string `json:"message"`
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
	return fmt.Sprintf(
		"HTTP Status Code [%d]: %s\nFortnox Code: %d <-> %s",
		f.HTTPStatus, http.StatusText(f.HTTPStatus), f.Code, f.Message)
}

// Translate translates error message to languages defined by [langCode]
func (f FortnoxError) Translate(langCode LangCode) string {
	langToDescriptionMapping := CodeToLanguagesMapping[f.Code]
	description := langToDescriptionMapping[langCode]
	return description
}

// ErrorMessage response type
type ErrorMessage struct {
	Error   int    `json:"Error"`
	Message string `json:"Message"`
	Code    int    `json:"Code"`
}

// Translate translates error message to languages defined by [langCode]
func (f ErrorMessage) Translate(langCode LangCode) string {
	langToDescriptionMapping := CodeToLanguagesMapping[f.Code]
	description := langToDescriptionMapping[langCode]
	return description
}

type LangCode string

const (
	EN LangCode = "EN"
	SE LangCode = "SE"
)

var CodeToLanguagesMapping = map[int]map[LangCode]string{
	1000003: {
		SE: "System exception",
		EN: "Something went wrong on our end, please contact us.",
	},
	1000030: {
		SE: "Invalid response type",
		EN: "The provided response type(Accept) was invalid.",
	},
	1000031: {
		SE: "Invalid content type",
		EN: "The provided content type was invalid.",
	},
	2000106: {
		SE: "Värdet måste vara alfanumeriskt ({Value})",
		EN: "The value needs to be alphanumeric.",
	},
	2000108: {
		SE: "Värdet måste vara numeriskt ({Value})",
		EN: "The value needs to be numeric.",
	},
	2000134: {
		SE: "Värdet måste vara en boolean ({Value})",
		EN: "The value needs to be boolean",
	},
	2000310: {
		SE: "Ogiltig inloggning",
		EN: "The Client-Secret or the Access-Token is either missing or is incorrect.",
	},
	2000311: {
		SE: "Kan inte logga in, access-token eller client-secret saknas(2).",
		EN: "The Client-Secret or the Access-Token is either missing or is incorrect.",
	},
	2000359: {
		SE: "Värdet innehåller ej tillåtna tecken. ({Value})",
		EN: "The value contains invalid characters.",
	},
	2000588: {
		SE: "Ogiltig parameter i anropet.",
		EN: "A parameter is invalid. Read more about parameters. [https://developer.fortnox.se/documentation/general/parameters/]",
	},
	2000637: {
		SE: "Kundnummer 1 används redan. Kundnumret har redan använts men blivit raderat.",
		EN: "Customer number 1 is/has already been used.",
	},
	2000729: {
		SE: "A valid identifier was not provided.",
		EN: "A valid identifier was not provided.",
	},
	2001103: {
		SE: "Api-licens saknas.",
		EN: "The requested Fortnox account does not have a license to use the API",
	},
	2001392: {
		SE: "Ingen eller felaktig typ av data.",
		EN: "The request body was empty or contained incorrect data.",
	},
	2001740: {
		SE: "Inläsning av dokument misslyckades: {Message}",
		EN: "The XML object contained an error.",
	},
	2002115: {
		SE: "Error deserializing JSON: JSON_ERROR_SYNTAX",
		EN: "The JSON object contained an error.",
	},
	2001304: {
		SE: "Kunde inte hitta konto",
		EN: "Could not find Account",
	},
	2001399: {
		SE: "Felaktigt fältnamn",
		EN: "Invalid Field name",
	},
	2001101: {
		SE: "Det finns ingen aktiv licens för önskat scope",
		EN: "There is no active licens for the desired scope.",
	},
	2000663: {
		SE: "Har inte behörighet för scope",
		EN: "No access to the current scope.",
	},
	2003095: {
		SE: "Det saknas ett förvalt konto för Inköp SE, omvänd skattskyldighet\t",
		EN: "Account is missing for Purchase SE reversed tax liability",
	},
	2000755: {
		SE: "Leverantörsfakturan balanserar inte",
		EN: "Supplier invoice does not balance",
	},
	2003115: {
		SE: "Momsrader för momstyp REVERSE måste vara märkta med motsvarande CODE",
		EN: "Tax Rows for VAT type REVERSE must be marked with CODE",
	},
	2003124: {
		SE: "Enbart ordrar som levererats ut kan klarmarkeras",
		EN: "",
	},
	2003125: {
		SE: "Ett klarmarkerat dokument kan inte ändras",
		EN: "",
	},
	2003126: {
		SE: "Utleveransdatum kan inte vara senare än dagens datum",
		EN: "",
	},
	2003241: {
		SE: "Migrering är redan påbörjad eller avslutad",
		EN: "",
	},
	2003275: {
		SE: "Ej autentiserad",
		EN: "",
	},
	2003277: {
		SE: "Hittades inte i lagermodulen",
		EN: "",
	},
	2003399: {
		SE: "Dokumentet är makulerat i lagermodulen",
		EN: "",
	},
	2003127: {
		SE: "Ett fel uppstod i lagermodulen",
		EN: "",
	},
	2000204: {
		SE: "Kunde inte hämta/hitta kund (kundnummer)",
		EN: "The customer in the request is not available in the customer resource.",
	},
	2000433: {
		SE: "Kunde inte hämta/hitta kund (kundnummer)",
		EN: "The customer in the request is not available in the customer resource.",
	},
	2001302: {
		SE: "Kunde inte hitta artikel",
		EN: "Could not find article used in request",
	},
	2000428: {
		SE: "Kan inte hitta artikeln",
		EN: "",
	},
}
