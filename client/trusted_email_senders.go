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
	ctx context.Context) (*GetAllTrustedAndRejectedEmailSendersResp, error) {

	resp := &GetAllTrustedAndRejectedEmailSendersResp{}

	err := c._GET(ctx, emailSendersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateTrustedEmailAddress does _POST https://api.fortnox.se/3/emailsenders/trusted
//
// req - trusted email sender to create
func (c *Client) CreateTrustedEmailAddress(
	ctx context.Context,
	req *CreateTrustedEmailAddressReq) (*CreateTrustedEmailAddressResp, error) {

	resp := &CreateTrustedEmailAddressResp{}

	uri := fmt.Sprintf("%s/trusted", emailSendersURI)

	err := c._POST(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveTrustedEmailAddress does _DELETE https://api.fortnox.se/3/emailsenders/trusted/{Id}
//
// id - identifies the trusted email sender to delete
func (c *Client) RemoveTrustedEmailAddress(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/trusted/%d", unitsURI, id)
	return c._DELETE(ctx, uri)
}

type GetAllTrustedAndRejectedEmailSendersResp struct {
	EmailSenders struct {
		TrustedSenders []struct {
			Id    int    `json:"Id"`
			Email string `json:"Email"`
		} `json:"TrustedSenders"`
		RejectedSenders []struct {
			Id    int    `json:"Id"`
			Email string `json:"Email"`
		} `json:"RejectedSenders"`
	} `json:"EmailSenders"`
}

type CreateTrustedEmailAddressReq struct {
	TrustedSender struct {
		Id    int    `json:"Id"`
		Email string `json:"Email"`
	} `json:"TrustedSender"`
}

type CreateTrustedEmailAddressResp struct {
	TrustedSender struct {
		Id    int    `json:"Id"`
		Email string `json:"Email"`
	} `json:"TrustedSender"`
}
