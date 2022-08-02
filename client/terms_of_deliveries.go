package client

import (
	"context"
	"fmt"
)

const (
	termsOfDeliveriesURI = "termsofdeliveries"
)

// GetAllTermsOfDeliveries does _GET https://api.fortnox.se/3/termsofdeliveries
func (c *Client) GetAllTermsOfDeliveries(ctx context.Context) ([]TermsOfDelivery, error) {
	resp := &GetAllTermsOfDeliveriesResp{}

	err := c._GET(ctx, termsOfDeliveriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.TermsOfDeliveries, nil
}

// CreateTermsOfDeliveries does _POST https://api.fortnox.se/3/termsofdeliveries
//
// req - terms of delivery to create
func (c *Client) CreateTermsOfDeliveries(
	ctx context.Context,
	tod *TermsOfDelivery) (*TermsOfDelivery, error) {

	req := &CreateTermsOfDeliveriesReq{TermsOfDelivery: *tod}
	resp := &CreateTermsOfDeliveriesResp{}

	err := c._POST(ctx, termsOfDeliveriesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfDelivery, nil
}

// GetTermOfDelivery does _GET https://api.fortnox.se/3/termsofdeliveries/{Code}
//
// code - identifies the terms of delivery
func (c *Client) GetTermOfDelivery(ctx context.Context, code string) (*TermsOfDelivery, error) {
	resp := &GetTermOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", termsOfDeliveriesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfDelivery, nil
}

// UpdateTermOfDelivery does _PUT https://api.fortnox.se/3/termsofdeliveries/{Code}
//
// code - identifies the terms of delivery
//
// req - terms of delivery to update
func (c *Client) UpdateTermOfDelivery(
	ctx context.Context,
	code string,
	tod *TermsOfDelivery) (*TermsOfDelivery, error) {

	req := &UpdateTermOfDeliveryReq{TermsOfDelivery: *tod}
	resp := &UpdateTermOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", termsOfDeliveriesURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TermsOfDelivery, nil
}

type TermsOfDelivery struct {
	Url                string `json:"@url"`
	Code               string `json:"Code"`
	Description        string `json:"Description"`
	DescriptionEnglish string `json:"DescriptionEnglish"`
}

type GetAllTermsOfDeliveriesResp struct {
	TermsOfDeliveries []TermsOfDelivery `json:"TermsOfDeliveries"`
}

type CreateTermsOfDeliveriesReq struct {
	TermsOfDelivery TermsOfDelivery `json:"TermsOfDelivery"`
}

type CreateTermsOfDeliveriesResp struct {
	TermsOfDelivery TermsOfDelivery `json:"TermsOfDelivery"`
}

type GetTermOfDeliveryResp struct {
	TermsOfDelivery TermsOfDelivery `json:"TermsOfDelivery"`
}

type UpdateTermOfDeliveryReq struct {
	TermsOfDelivery TermsOfDelivery `json:"TermsOfDelivery"`
}

type UpdateTermOfDeliveryResp struct {
	TermsOfDelivery TermsOfDelivery `json:"TermsOfDelivery"`
}
