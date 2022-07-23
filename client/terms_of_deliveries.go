package client

import (
	"context"
	"fmt"
)

const (
	termsOfDeliveriesURI = "termsofdeliveries"
)

// GetAllTermsOfDeliveries does _GET https://api.fortnox.se/3/termsofdeliveries
func (c *Client) GetAllTermsOfDeliveries(ctx context.Context) (*GetAllTermsOfDeliveriesResp, error) {
	resp := &GetAllTermsOfDeliveriesResp{}

	err := c._GET(ctx, termsOfDeliveriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateTermsOfDeliveries does _POST https://api.fortnox.se/3/termsofdeliveries
//
// req - terms of delivery to create
func (c *Client) CreateTermsOfDeliveries(
	ctx context.Context,
	req *CreateTermsOfDeliveriesReq) (*CreateTermsOfDeliveriesResp, error) {

	resp := &CreateTermsOfDeliveriesResp{}

	err := c._POST(ctx, termsOfDeliveriesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTermOfDelivery does _GET https://api.fortnox.se/3/termsofdeliveries/{Code}
//
// code - identifies the terms of delivery
func (c *Client) GetTermOfDelivery(ctx context.Context, code string) (*GetTermOfDeliveryResp, error) {
	resp := &GetTermOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", termsOfDeliveriesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateTermOfDelivery does _PUT https://api.fortnox.se/3/termsofdeliveries/{Code}
//
// code - identifies the terms of delivery
//
// req - terms of delivery to update
func (c *Client) UpdateTermOfDelivery(
	ctx context.Context,
	code string, req *UpdateTermOfDeliveryReq) (*UpdateTermOfDeliveryResp, error) {

	resp := &UpdateTermOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", termsOfDeliveriesURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllTermsOfDeliveriesResp struct {
	TermsOfDeliveries []struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDeliveries"`
}

type CreateTermsOfDeliveriesReq struct {
	TermsOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDelivery"`
}

type CreateTermsOfDeliveriesResp struct {
	TermsOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDelivery"`
}

type GetTermOfDeliveryResp struct {
	TermsOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDelivery"`
}

type UpdateTermOfDeliveryReq struct {
	TermsOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDelivery"`
}

type UpdateTermOfDeliveryResp struct {
	TermsOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"TermsOfDelivery"`
}
