package client

import (
	"context"
	"fmt"
)

const (
	customerReferencesURI = "customerreferences"
)

// GetAllCustomerReferences does _GET https://api.fortnox.se/3/customerreferences/
func (c *Client) GetAllCustomerReferences(ctx context.Context) (*CustomerReference, error) {
	resp := &GetAllCustomerReferencesResp{}

	err := c._GET(ctx, customerReferencesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CustomerReference, nil
}

// CreateCustomerReference does _POST https://api.fortnox.se/3/customerreferences/
//
// crr - customer reference row to create
func (c *Client) CreateCustomerReference(ctx context.Context, crr *CustomerReferenceRow) (*CustomerReference, error) {
	req := &CreateCustomerReferenceReq{CustomerReferenceRow: *crr}
	resp := &CreateCustomerReferenceResp{}

	err := c._POST(ctx, customerReferencesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.CustomerReference, nil
}

// GetCustomerReference does _GET https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
func (c *Client) GetCustomerReference(
	ctx context.Context,
	customerReferenceRowID string) (*GetCustomerReferenceResp, error) {

	resp := &GetCustomerReferenceResp{}

	uri := fmt.Sprintf("%s/%s", customerReferencesURI, customerReferenceRowID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateCustomerReference does _PUT https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
//
// crr - customer reference row to update
func (c *Client) UpdateCustomerReference(
	ctx context.Context,
	customerReferenceRowID string,
	crr *CustomerReferenceRow) (*Customer, error) {

	req := &UpdateCustomerReferenceReq{CustomerReferenceRow: *crr}
	resp := &UpdateCustomerReferenceResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerReferenceRowID)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// DeleteCustomerReferenceRow does _DELETE https://api.fortnox.se/3/customerreferences/{CustomerReferenceRowId}
//
// customerReferenceRowID - identifies the customer reference row
func (c *Client) DeleteCustomerReferenceRow(ctx context.Context, customerReferenceRowID string) error {
	uri := fmt.Sprintf("%s/%s", customerReferencesURI, customerReferenceRowID)
	return c._DELETE(ctx, uri)
}

type GetAllCustomerReferencesResp struct {
	CustomerReference CustomerReference `json:"CustomerReference"`
}

type CreateCustomerReferenceReq struct {
	CustomerReferenceRow CustomerReferenceRow `json:"CustomerReferenceRow"`
}

type CreateCustomerReferenceResp struct {
	CustomerReference CustomerReference `json:"CustomerReference"`
}

type GetCustomerReferenceResp struct {
	CustomerReference CustomerReference `json:"CustomerReference"`
}

type CustomerReference struct {
	CustomerReferenceRows []CustomerReferenceRow `json:"CustomerReferenceRows"`
}

type UpdateCustomerReferenceReq struct {
	CustomerReferenceRow CustomerReferenceRow `json:"CustomerReferenceRow"`
}

type CustomerReferenceRow struct {
	Reference      string `json:"Reference,omitempty"`
	CustomerNumber int64  `json:"CustomerNumber,omitempty"`
	Id             int    `json:"Id,omitempty"`
}

type UpdateCustomerReferenceResp struct {
	Customer Customer `json:"Customer"`
}
