package client

import (
	"context"
	"fmt"
)

const (
	wayOfDeliveriesURI = "wayofdeliveries"
)

// GetAllWayOfDeliveries does _GET https://api.fortnox.se/3/wayofdeliveries
func (c *Client) GetAllWayOfDeliveries(ctx context.Context) (*GetAllWayOfDeliveriesResp, error) {
	resp := &GetAllWayOfDeliveriesResp{}

	err := c._GET(ctx, wayOfDeliveriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateWayOfDeliveries does _POST https://api.fortnox.se/3/wayofdeliveries
//
// req - way of delivery to create
func (c *Client) CreateWayOfDeliveries(
	ctx context.Context,
	req *CreateWayOfDeliveriesReq) (*CreateWayOfDeliveriesResp, error) {

	resp := &CreateWayOfDeliveriesResp{}

	err := c._POST(ctx, wayOfDeliveriesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetWayOfDeliveryByCode does _GET https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
func (c *Client) GetWayOfDeliveryByCode(ctx context.Context, code string) (*GetWayOfDeliveryByCodeResp, error) {
	resp := &GetWayOfDeliveryByCodeResp{}

	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateWayOfDelivery does UPDATE https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
//
// req - way of delivery to update
func (c *Client) UpdateWayOfDelivery(
	ctx context.Context,
	code string,
	req *UpdateWayOfDeliveryReq) (*UpdateWayOfDeliveryResp, error) {

	resp := &UpdateWayOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)

	err := c._POST(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveWayOfDelivery does _DELETE https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
func (c *Client) RemoveWayOfDelivery(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)
	return c._DELETE(ctx, uri)
}

type GetAllWayOfDeliveriesResp struct {
	WayOfDeliveries []struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDeliveries"`
}

type CreateWayOfDeliveriesReq struct {
	WayOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDelivery"`
}

type CreateWayOfDeliveriesResp struct {
	WayOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDelivery"`
}

type GetWayOfDeliveryByCodeResp struct {
	WayOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDelivery"`
}

type UpdateWayOfDeliveryReq struct {
	WayOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDelivery"`
}

type UpdateWayOfDeliveryResp struct {
	WayOfDelivery struct {
		Url                string `json:"@url"`
		Code               string `json:"Code"`
		Description        string `json:"Description"`
		DescriptionEnglish string `json:"DescriptionEnglish"`
	} `json:"WayOfDelivery"`
}
