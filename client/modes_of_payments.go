package client

import (
	"context"
	"fmt"
)

const (
	modesOfPaymentsURI = "modesofpayments"
)

// GetAllModesOfPayments does _GET https://api.fortnox.se/3/modesofpayments
func (c *Client) GetAllModesOfPayments(ctx context.Context) ([]ModeOfPayment, error) {
	resp := &GetAllModesOfPaymentsResp{}

	err := c._GET(ctx, modesOfPaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.ModeOfPayments, nil
}

// CreateModeOfPayment does _POST https://api.fortnox.se/3/modesofpayments
//
// mop - mode of payment to create
func (c *Client) CreateModeOfPayment(ctx context.Context, si *ModeOfPayment) (*ModeOfPayment, error) {
	req := &CreateModeOfPaymentReq{ModeOfPayment: *si}
	resp := &CreateModeOfPaymentResp{}

	err := c._POST(ctx, modesOfPaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ModeOfPayment, nil
}

// GetModeOfPayment does _GET https://api.fortnox.se/3/modesofpayments/{Code}
//
// code - identifies the mode of payment
func (c *Client) GetModeOfPayment(ctx context.Context, code string) (*ModeOfPayment, error) {
	resp := &GetModeOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", modesOfPaymentsURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ModeOfPayment, nil
}

// UpdateModeOfPayment does _PUT https://api.fortnox.se/3/modesofpayments/{Code}
//
// code - identifies the mode of payment
//
// mop - mode of payment to update
func (c *Client) UpdateModeOfPayment(
	ctx context.Context,
	code string,
	mop *ModeOfPayment) (*ModeOfPayment, error) {

	req := &UpdateModeOfPaymentReq{ModeOfPayment: *mop}
	resp := &UpdateModeOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", modesOfPaymentsURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ModeOfPayment, nil
}

type GetAllModesOfPaymentsResp struct {
	ModeOfPayments []ModeOfPayment `json:"ModeOfPayments"`
}

type CreateModeOfPaymentReq struct {
	ModeOfPayment ModeOfPayment `json:"ModeOfPayment"`
}

type CreateModeOfPaymentResp struct {
	ModeOfPayment ModeOfPayment `json:"ModeOfPayment"`
}

type GetModeOfPaymentResp struct {
	ModeOfPayment ModeOfPayment `json:"ModeOfPayment"`
}

type UpdateModeOfPaymentReq struct {
	ModeOfPayment ModeOfPayment `json:"ModeOfPayment"`
}

type UpdateModeOfPaymentResp struct {
	ModeOfPayment ModeOfPayment `json:"ModeOfPayment"`
}
type ModeOfPayment struct {
	Url           string `json:"@url,omitempty"`
	Code          string `json:"Code,omitempty"`
	Description   string `json:"Description,omitempty"`
	AccountNumber string `json:"AccountNumber,omitempty"`
}
