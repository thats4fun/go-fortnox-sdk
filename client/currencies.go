package client

import (
	"context"
	"fmt"
)

const (
	currenciesURI = "currencies"
)

// GetAllCurrencies does _GET https://api.fortnox.se/3/currencies
func (c *Client) GetAllCurrencies(ctx context.Context) ([]Currency, error) {
	resp := &GetAllCurrenciesResp{}

	err := c._GET(ctx, currenciesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Currencies, nil
}

// CreateCurrency does _POST https://api.fortnox.se/3/currencies
//
// req - currency to create
func (c *Client) CreateCurrency(ctx context.Context, cur *Currency) (*Currency, error) {
	req := &CreateCurrencyReq{Currency: *cur}
	resp := &CreateCurrencyResp{}

	err := c._POST(ctx, currenciesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Currency, nil
}

// GetCurrency does _GET https://api.fortnox.se/3/currencies/{Code}
//
// code - identifies currency
func (c *Client) GetCurrency(ctx context.Context, code string) (*Currency, error) {
	resp := &GetCurrencyResp{}

	uri := fmt.Sprintf("%s/%s", currenciesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Currency, nil
}

// UpdateCurrency does _PUT https://api.fortnox.se/3/currencies/{Code}
//
// code - identifies currency
//
// req - currency to update
func (c *Client) UpdateCurrency(ctx context.Context, code string, cur *Currency) (*Currency, error) {
	req := &UpdateCurrencyReq{Currency: *cur}
	resp := &UpdateCurrencyResp{}

	uri := fmt.Sprintf("%s/%s", currenciesURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Currency, nil
}

// RemoveCurrency does _DELETE
//
// code - identifies the currency to remove
func (c *Client) RemoveCurrency(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", currenciesURI, code)
	return c._DELETE(ctx, uri)
}

type Currency struct {
	Currency string  `json:"currency,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Unit     int     `json:"unit,omitempty"`
}

type GetAllCurrenciesResp struct {
	Currencies []Currency `json:"Currencies"`
}

type CreateCurrencyReq struct {
	Currency Currency `json:"Currency"`
}

type CreateCurrencyResp struct {
	Currency Currency `json:"Currency"`
}

type GetCurrencyResp struct {
	Currency Currency `json:"Currency"`
}

type UpdateCurrencyReq struct {
	Currency Currency `json:"Currency"`
}

type UpdateCurrencyResp struct {
	Currency Currency `json:"Currency"`
}
