package client

import (
	"context"
	"fmt"
)

const (
	predefinedAccountsURI = "predefinedaccounts"
)

// GetAllPredefinedAccounts does _GET https://api.fortnox.se/3/predefinedaccounts/
func (c *Client) GetAllPredefinedAccounts(ctx context.Context) (*GetAllPredefinedAccountsResp, error) {
	resp := &GetAllPredefinedAccountsResp{}

	err := c._GET(ctx, predefinedAccountsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPredefinedAccount does _GET https://api.fortnox.se/3/predefinedaccounts/{name}
//
// name - identifies the predefined account
func (c *Client) GetPredefinedAccount(ctx context.Context, name string) (*GetPredefinedAccountResp, error) {
	resp := &GetPredefinedAccountResp{}

	uri := fmt.Sprintf("%s/%s", predefinedAccountsURI, name)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdatePredefinedAccount does _PUT https://api.fortnox.se/3/predefinedaccounts/{name}
//
// name - identifies the predefined account
//
// req - predefined account to update
func (c *Client) UpdatePredefinedAccount(
	ctx context.Context,
	name string,
	req *UpdatePredefinedAccountReq) (*UpdatePredefinedAccountResp, error) {

	resp := &UpdatePredefinedAccountResp{}

	uri := fmt.Sprintf("%s/%s", predefinedAccountsURI, name)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllPredefinedAccountsResp struct {
	PreDefinedAccounts []struct {
		Url     string `json:"@url"`
		Name    string `json:"Name"`
		Account int    `json:"Account"`
	} `json:"PreDefinedAccounts"`
}

type GetPredefinedAccountResp struct {
	PreDefinedAccount struct {
		Url     string `json:"@url"`
		Name    string `json:"Name"`
		Account int    `json:"Account"`
	} `json:"PreDefinedAccount"`
}

type UpdatePredefinedAccountReq struct {
	PreDefinedAccount struct {
		Url     string `json:"@url"`
		Name    string `json:"Name"`
		Account int    `json:"Account"`
	} `json:"PreDefinedAccount"`
}

type UpdatePredefinedAccountResp struct {
	PreDefinedAccount struct {
		Url     string `json:"@url"`
		Name    string `json:"Name"`
		Account int    `json:"Account"`
	} `json:"PreDefinedAccount"`
}
