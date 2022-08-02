package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// URI
const (
	accountsURI = "accounts"
)

// url query param names
const (
	financialYearParamName = "financialyear"

	lastModifiedParamName = "lastmodified"
	sortByParamName       = "sortby"
	sruParamName          = "sru"
)

// GetAccount does _GET https://api.fortnox.se/3/accounts/{Number}
//
// accountID - identifies the account
func (c *Client) GetAccount(ctx context.Context, accountID int) (*Account, error) {
	resp := &GetAccountResp{}

	uri := fmt.Sprintf("%s/%d", accountsURI, accountID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Account, nil
}

// UpdateAccount does _PUT https://api.fortnox.se/3/accounts/{Number}
//
// accountID - identifies the account
//
// financialYear - financial year to update account against, param is optional
func (c *Client) UpdateAccount(
	ctx context.Context,
	accountID int,
	a *Account,
	filter *FinancialYearFilter) (*Account, error) {

	req := &UpdateAccountReq{Account: *a}
	resp := &UpdateAccountResp{}

	uri := fmt.Sprintf("%s/%d", accountsURI, accountID)

	var params url.Values
	if filter != nil {
		params = filter.urlValues()
	}

	err := c._PUT(ctx, uri, params, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Account, nil
}

// GetAllAccounts does _GET https://api.fortnox.se/3/accounts/
//
// filter - GetAllAccountsFilter
func (c *Client) GetAllAccounts(ctx context.Context, filter *GetAllAccountsFilter) ([]Account, error) {
	resp := &GetAllAccountsResp{}

	var params url.Values
	if filter != nil {
		params = filter.urlValues()
	}

	err := c._GET(ctx, accountsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.Accounts, nil
}

// CreateAccount does _POST https://api.fortnox.se/3/accounts/
//
// filter - financial year to create account against
func (c *Client) CreateAccount(
	ctx context.Context,
	a *Account,
	filter *FinancialYearFilter) (*Account, error) {

	req := &CreateAccountReq{Account: *a}
	resp := CreateAccountResp{}

	var params url.Values
	if filter != nil {
		params = filter.urlValues()
	}

	err := c._POST(ctx, accountsURI, params, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Account, nil
}

type GetAllAccountsFilter struct {
	LastModified string
	SortBy       string
	SRU          int
}

func (f *GetAllAccountsFilter) urlValues() url.Values {
	params := url.Values{}

	if strings.TrimSpace(f.LastModified) != "" {
		params[lastModifiedParamName] = []string{f.LastModified}
	}

	if strings.TrimSpace(f.SortBy) != "" {
		params[sortByParamName] = []string{f.SortBy}
	}

	if f.SRU > 0 {
		params[sruParamName] = []string{strconv.Itoa(f.SRU)}
	}

	return params
}

type FinancialYearFilter struct {
	FinancialYear int
}

func (f *FinancialYearFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	if f.FinancialYear > 0 {
		params[financialYearParamName] = []string{strconv.Itoa(f.FinancialYear)}
	}

	return params
}

type Account struct {
	Url                            string            `json:"@url,omitempty"`
	Active                         bool              `json:"Active,omitempty"`
	BalanceBroughtForward          int               `json:"BalanceBroughtForward,omitempty"`
	CostCenter                     string            `json:"CostCenter,omitempty"`
	CostCenterSettings             string            `json:"CostCenterSettings,omitempty"`
	Description                    string            `json:"Description,omitempty"`
	Number                         int               `json:"Number,omitempty"`
	Project                        string            `json:"Project,omitempty"`
	ProjectSettings                string            `json:"ProjectSettings,omitempty"`
	SRU                            int               `json:"SRU,omitempty"`
	Year                           int               `json:"Year,omitempty"`
	VATCode                        string            `json:"VATCode,omitempty"`
	BalanceCarriedForward          int               `json:"BalanceCarriedForward,omitempty"`
	TransactionInformation         string            `json:"TransactionInformation,omitempty"`
	TransactionInformationSettings string            `json:"TransactionInformationSettings,omitempty"`
	QuantitySettings               string            `json:"QuantitySettings,omitempty"`
	QuantityUnit                   string            `json:"QuantityUnit,omitempty"`
	OpeningQuantities              []OpeningQuantity `json:"OpeningQuantities,omitempty"`
}

type OpeningQuantity struct {
	Project string `json:"Project,omitempty"`
	Balance int    `json:"Balance,omitempty"`
}

type GetAccountResp struct {
	Account Account `json:"Account"`
}

type UpdateAccountReq GetAccountResp

type UpdateAccountResp GetAccountResp

type GetAllAccountsResp struct {
	Accounts []Account `json:"Accounts"`
}

type CreateAccountReq struct {
	Account Account `json:"Account"`
}

type CreateAccountResp struct {
	Account Account `json:"Account"`
}
