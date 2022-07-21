package client

import (
	"context"
	"fmt"
)

const (
	priceListURI = "pricelist"
)

// GetAllPriceLists does _GET
func (c *Client) GetAllPriceLists(ctx context.Context) (*GetAllPriceListsResp, error) {
	resp := &GetAllPriceListsResp{}

	err := c._GET(ctx, priceListURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreatePriceList does _POST
// req -
func (c *Client) CreatePriceList(ctx context.Context, req *CreatePriceListReq) (*CreatePriceListResp, error) {
	resp := &CreatePriceListResp{}

	err := c._POST(ctx, priceListURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPriceList does _GET
// code -
func (c *Client) GetPriceList(ctx context.Context, code string) (*GetPriceListResp, error) {
	resp := &GetPriceListResp{}

	uri := fmt.Sprintf("%s/%s", priceListURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdatePriceList doe _PUT
// code -
// req -
func (c *Client) UpdatePriceList(
	ctx context.Context,
	code string,
	req *UpdatePriceListReq) (*UpdatePriceListResp, error) {

	resp := &UpdatePriceListResp{}

	uri := fmt.Sprintf("%s/%s", priceListURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllPriceListsResp struct {
	PriceLists []struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceLists"`
}

type CreatePriceListReq struct {
	PriceList struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceList"`
}

type CreatePriceListResp struct {
	PriceList struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceList"`
}

type GetPriceListResp struct {
	PriceList struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceList"`
}

type UpdatePriceListReq struct {
	PriceList struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceList"`
}

type UpdatePriceListResp struct {
	PriceList struct {
		Url         string `json:"@url"`
		Code        string `json:"Code"`
		Description string `json:"Description"`
		Comments    string `json:"Comments"`
		PreSelected bool   `json:"PreSelected"`
	} `json:"PriceList"`
}
