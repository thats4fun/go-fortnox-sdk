package client

import (
	"context"
	"fmt"
)

const (
	predefinedAccountsURI = "predefinedaccounts"
)

// GetAllPredefinedAccounts does _GET https://api.fortnox.se/3/predefinedaccounts/
func (c *Client) GetAllPredefinedAccounts(ctx context.Context) ([]PreDefinedAccount, error) {
	resp := &GetAllPredefinedAccountsResp{}

	err := c._GET(ctx, predefinedAccountsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.PreDefinedAccounts, nil
}

// GetPredefinedAccount does _GET https://api.fortnox.se/3/predefinedaccounts/{name}
//
// name - identifies the predefined account
func (c *Client) GetPredefinedAccount(ctx context.Context, name string) (*PreDefinedAccount, error) {
	resp := &GetPredefinedAccountResp{}

	uri := fmt.Sprintf("%s/%s", predefinedAccountsURI, name)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PreDefinedAccount, nil
}

// UpdatePredefinedAccount does _PUT https://api.fortnox.se/3/predefinedaccounts/{name}
//
// name - identifies the predefined account
//
// pda - predefined account to update
func (c *Client) UpdatePredefinedAccount(
	ctx context.Context,
	name string,
	pda *PreDefinedAccount) (*PreDefinedAccount, error) {

	req := &UpdatePredefinedAccountReq{PreDefinedAccount: *pda}
	resp := &UpdatePredefinedAccountResp{}

	uri := fmt.Sprintf("%s/%s", predefinedAccountsURI, name)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.PreDefinedAccount, nil
}

type PreDefinedAccount struct {
	Url     string `json:"@url,omitempty"`
	Name    string `json:"Name,omitempty"`
	Account int    `json:"Account,omitempty"`
}

type GetAllPredefinedAccountsResp struct {
	PreDefinedAccounts []PreDefinedAccount `json:"PreDefinedAccounts"`
}

type GetPredefinedAccountResp struct {
	PreDefinedAccount PreDefinedAccount `json:"PreDefinedAccount"`
}

type UpdatePredefinedAccountReq struct {
	PreDefinedAccount PreDefinedAccount `json:"PreDefinedAccount"`
}

type UpdatePredefinedAccountResp struct {
	PreDefinedAccount PreDefinedAccount `json:"PreDefinedAccount"`
}
