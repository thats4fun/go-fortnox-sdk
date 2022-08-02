package client

import (
	"context"
	"fmt"
)

const (
	priceListURI = "pricelist"
)

// GetAllPriceLists does _GET https://api.fortnox.se/3/pricelists
func (c *Client) GetAllPriceLists(ctx context.Context) ([]PriceList, error) {
	resp := &GetAllPriceListsResp{}

	err := c._GET(ctx, priceListURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.PriceLists, nil
}

// CreatePriceList does _POST https://api.fortnox.se/3/pricelists
//
// req - price list to create
func (c *Client) CreatePriceList(ctx context.Context, pl *PriceList) (*PriceList, error) {
	req := CreatePriceListReq{PriceList: *pl}
	resp := &CreatePriceListResp{}

	err := c._POST(ctx, priceListURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PriceList, nil
}

// GetPriceList does _GET https://api.fortnox.se/3/pricelists/{Code}
//
// code - identifies the price list
func (c *Client) GetPriceList(ctx context.Context, code string) (*PriceList, error) {
	resp := &GetPriceListResp{}

	uri := fmt.Sprintf("%s/%s", priceListURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PriceList, nil
}

// UpdatePriceList does _PUT https://api.fortnox.se/3/pricelists/{Code}
//
// code - identifies the price list
//
// req - price list to update
func (c *Client) UpdatePriceList(
	ctx context.Context,
	code string,
	pl *PriceList) (*PriceList, error) {

	req := UpdatePriceListReq{PriceList: *pl}
	resp := &UpdatePriceListResp{}

	uri := fmt.Sprintf("%s/%s", priceListURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PriceList, nil
}

type PriceList struct {
	Url         string `json:"@url"`
	Code        string `json:"Code"`
	Description string `json:"Description"`
	Comments    string `json:"Comments"`
	PreSelected bool   `json:"PreSelected"`
}

type GetAllPriceListsResp struct {
	PriceLists []PriceList `json:"PriceLists"`
}

type CreatePriceListReq struct {
	PriceList PriceList `json:"PriceList"`
}

type CreatePriceListResp struct {
	PriceList PriceList `json:"PriceList"`
}

type GetPriceListResp struct {
	PriceList PriceList `json:"PriceList"`
}

type UpdatePriceListReq struct {
	PriceList PriceList `json:"PriceList"`
}

type UpdatePriceListResp struct {
	PriceList PriceList `json:"PriceList"`
}
