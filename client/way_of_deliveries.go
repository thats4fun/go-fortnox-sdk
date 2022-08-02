package client

import (
	"context"
	"fmt"
)

const (
	wayOfDeliveriesURI = "wayofdeliveries"
)

// GetAllWayOfDeliveries does _GET https://api.fortnox.se/3/wayofdeliveries
func (c *Client) GetAllWayOfDeliveries(ctx context.Context) ([]WayOfDelivery, error) {
	resp := &GetAllWayOfDeliveriesResp{}

	err := c._GET(ctx, wayOfDeliveriesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.WayOfDeliveries, nil
}

// CreateWayOfDeliveries does _POST https://api.fortnox.se/3/wayofdeliveries
//
// req - way of delivery to create
func (c *Client) CreateWayOfDeliveries(
	ctx context.Context,
	wod *WayOfDelivery) (*WayOfDelivery, error) {

	req := CreateWayOfDeliveriesReq{WayOfDelivery: *wod}
	resp := &CreateWayOfDeliveriesResp{}

	err := c._POST(ctx, wayOfDeliveriesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.WayOfDelivery, nil
}

// GetWayOfDeliveryByCode does _GET https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
func (c *Client) GetWayOfDeliveryByCode(ctx context.Context, code string) (*WayOfDelivery, error) {
	resp := &GetWayOfDeliveryByCodeResp{}

	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.WayOfDelivery, nil
}

// UpdateWayOfDelivery does _PUT https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
//
// req - way of delivery to update
func (c *Client) UpdateWayOfDelivery(
	ctx context.Context,
	code string,
	wod *WayOfDelivery) (*WayOfDelivery, error) {

	req := UpdateWayOfDeliveryReq{WayOfDelivery: *wod}
	resp := &UpdateWayOfDeliveryResp{}

	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.WayOfDelivery, nil
}

// RemoveWayOfDelivery does _DELETE https://api.fortnox.se/3/wayofdeliveries/{Code}
//
// code - identifies the way of delivery
func (c *Client) RemoveWayOfDelivery(ctx context.Context, code string) error {
	uri := fmt.Sprintf("%s/%s", wayOfDeliveriesURI, code)
	return c._DELETE(ctx, uri)
}

type WayOfDelivery struct {
	Url                string `json:"@url"`
	Code               string `json:"Code"`
	Description        string `json:"Description"`
	DescriptionEnglish string `json:"DescriptionEnglish"`
}

type GetAllWayOfDeliveriesResp struct {
	WayOfDeliveries []WayOfDelivery `json:"WayOfDeliveries"`
}

type CreateWayOfDeliveriesReq struct {
	WayOfDelivery WayOfDelivery `json:"WayOfDelivery"`
}

type CreateWayOfDeliveriesResp struct {
	WayOfDelivery WayOfDelivery `json:"WayOfDelivery"`
}

type GetWayOfDeliveryByCodeResp struct {
	WayOfDelivery WayOfDelivery `json:"WayOfDelivery"`
}

type UpdateWayOfDeliveryReq struct {
	WayOfDelivery WayOfDelivery `json:"WayOfDelivery"`
}

type UpdateWayOfDeliveryResp struct {
	WayOfDelivery WayOfDelivery `json:"WayOfDelivery"`
}
