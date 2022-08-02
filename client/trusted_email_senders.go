package client

import (
	"context"
	"fmt"
)

const (
	emailSendersURI = "emailsenders"
)

// GetAllTrustedAndRejectedEmailSenders does _GET https://api.fortnox.se/3/emailsenders/
func (c *Client) GetAllTrustedAndRejectedEmailSenders(
	ctx context.Context) (*EmailSenders, error) {

	resp := &GetAllTrustedAndRejectedEmailSendersResp{}

	err := c._GET(ctx, emailSendersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.EmailSenders, nil
}

// CreateTrustedEmailAddress does _POST https://api.fortnox.se/3/emailsenders/trusted
//
// req - trusted email sender to create
func (c *Client) CreateTrustedEmailAddress(
	ctx context.Context,
	ts *TrustedSender) (*TrustedSender, error) {

	req := &CreateTrustedEmailAddressReq{TrustedSender: *ts}
	resp := &CreateTrustedEmailAddressResp{}

	uri := fmt.Sprintf("%s/trusted", emailSendersURI)

	err := c._POST(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TrustedSender, nil
}

// RemoveTrustedEmailAddress does _DELETE https://api.fortnox.se/3/emailsenders/trusted/{Id}
//
// id - identifies the trusted email sender to delete
func (c *Client) RemoveTrustedEmailAddress(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/trusted/%d", unitsURI, id)
	return c._DELETE(ctx, uri)
}

type EmailSenders struct {
	TrustedSenders  []TrustedSender   `json:"TrustedSenders"`
	RejectedSenders []RejectedSenders `json:"RejectedSenders"`
}

type TrustedSender struct {
	Id    int    `json:"Id"`
	Email string `json:"Email"`
}

type RejectedSenders struct {
	Id    int    `json:"Id"`
	Email string `json:"Email"`
}

type GetAllTrustedAndRejectedEmailSendersResp struct {
	EmailSenders EmailSenders `json:"EmailSenders"`
}

type CreateTrustedEmailAddressReq struct {
	TrustedSender TrustedSender `json:"TrustedSender"`
}

type CreateTrustedEmailAddressResp struct {
	TrustedSender TrustedSender `json:"TrustedSender"`
}
