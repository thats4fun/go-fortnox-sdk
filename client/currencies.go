package client

import (
	"context"
	"fmt"
)

const (
	currenciesURI = "currencies"
)

// GetAllCurrencies does _GET https://api.fortnox.se/3/currencies
func (c *Client) GetAllCurrencies(ctx context.Context) (*GetAllCurrenciesResp, error) {
	resp := &GetAllCurrenciesResp{}

	err := c._GET(ctx, currenciesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateCurrency does _POST https://api.fortnox.se/3/currencies
//
// req - currency to create
func (c *Client) CreateCurrency(ctx context.Context, req *CreateCurrencyReq) (*CreateCurrencyResp, error) {
	resp := &CreateCurrencyResp{}

	err := c._POST(ctx, currenciesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCurrency does _GET https://api.fortnox.se/3/currencies/{Code}
//
// code - identifies currency
func (c *Client) GetCurrency(ctx context.Context, code string) (*GetCurrencyResp, error) {
	resp := &GetCurrencyResp{}

	uri := fmt.Sprintf("%s/%s", currenciesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateCurrency does _PUT https://api.fortnox.se/3/currencies/{Code}
//
// code - identifies currency
//
// req - currency to update
func (c *Client) UpdateCurrency(ctx context.Context, code string, req *UpdateCurrencyReq) (*UpdateCurrencyResp, error) {
	resp := &UpdateCurrencyResp{}

	uri := fmt.Sprintf("%s/%s", currenciesURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveCurrency does _DELETE
//
// code - identifies the currency to remove
func (c *Client) RemoveCurrency(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", currenciesURI, code)
	return c._DELETE(ctx, uri)
}

type GetAllCurrenciesResp struct {
	Currencies []struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currencies"`
}

type CreateCurrencyReq struct {
	Currency struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currency"`
}

type CreateCurrencyResp struct {
	Currency struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currency"`
}

type GetCurrencyResp struct {
	Currency struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currency"`
}

type UpdateCurrencyReq struct {
	Currency struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currency"`
}

type UpdateCurrencyResp struct {
	Currency struct {
		Currency string  `json:"currency"`
		Rate     float64 `json:"rate"`
		Unit     int     `json:"unit"`
	} `json:"Currency"`
}
