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
func (c *Client) GetAccount(ctx context.Context, accountID int) (*GetAccountResp, error) {
	resp := &GetAccountResp{}

	uri := fmt.Sprintf("%s/%d", accountsURI, accountID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateAccount does _PUT https://api.fortnox.se/3/accounts/{Number}
//
// accountID - identifies the account
//
// financialYear - financial year to update account against, param is optional
func (c *Client) UpdateAccount(
	ctx context.Context,
	accountID int,
	req *UpdateAccountReq,
	filter *FinancialYearFilter) (*UpdateAccountResp, error) {

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

	return resp, nil
}

// GetAllAccounts does _GET https://api.fortnox.se/3/accounts/
//
// filter - GetAllAccountsFilter
func (c *Client) GetAllAccounts(ctx context.Context, filter *GetAllAccountsFilter) ([]*AccountResult, error) {
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
	req *CreateAccountReq,
	filter *FinancialYearFilter) (*CreateAccountInnerResp, error) {

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

func (f FinancialYearFilter) urlValues() url.Values {
	params := url.Values{}

	if f.FinancialYear > 0 {
		params[financialYearParamName] = []string{strconv.Itoa(f.FinancialYear)}
	}

	return params
}

type OpeningQuantity struct {
	Project string `json:"Project"`
	Balance int    `json:"Balance"`
}

type GetAccountResp struct {
	Account struct {
		Url                            string `json:"@url"`
		Active                         bool   `json:"Active"`
		BalanceBroughtForward          int    `json:"BalanceBroughtForward"`
		CostCenter                     string `json:"CostCenter"`
		CostCenterSettings             string `json:"CostCenterSettings"`
		Description                    string `json:"Description"`
		Number                         int    `json:"Number"`
		Project                        string `json:"Project"`
		ProjectSettings                string `json:"ProjectSettings"`
		SRU                            int    `json:"SRU"`
		Year                           int    `json:"Year"`
		VATCode                        string `json:"VATCode"`
		BalanceCarriedForward          int    `json:"BalanceCarriedForward"`
		TransactionInformation         string `json:"TransactionInformation"`
		TransactionInformationSettings string `json:"TransactionInformationSettings"`
		QuantitySettings               string `json:"QuantitySettings"`
		QuantityUnit                   string `json:"QuantityUnit"`
		OpeningQuantities              []struct {
			Project string `json:"Project"`
			Balance int    `json:"Balance"`
		} `json:"OpeningQuantities"`
	} `json:"Account"`
}

type UpdateAccountReq GetAccountResp

type UpdateAccountResp GetAccountResp

type GetAllAccountsResp struct {
	Accounts []*AccountResult `json:"Accounts"`
}

type AccountResult struct {
	Url                   string `json:"@url"`
	Active                bool   `json:"Active"`
	BalanceBroughtForward int    `json:"BalanceBroughtForward"`
	CostCenter            string `json:"CostCenter"`
	CostCenterSettings    string `json:"CostCenterSettings"`
	Description           string `json:"Description"`
	Number                int    `json:"Number"`
	Project               string `json:"Project"`
	ProjectSettings       string `json:"ProjectSettings"`
	SRU                   int    `json:"SRU"`
	Year                  int    `json:"Year"`
	VATCode               string `json:"VATCode"`
}

type CreateAccountReq struct {
	Account CreateAccountInnerReq `json:"Account"`
}

type CreateAccountInnerReq struct {
	Active                         bool              `json:"Active"`
	BalanceBroughtForward          int               `json:"BalanceBroughtForward"`
	CostCenter                     string            `json:"CostCenter"`
	CostCenterSettings             string            `json:"CostCenterSettings"`
	Description                    string            `json:"Description"`
	Number                         int               `json:"Number"`
	Project                        string            `json:"Project"`
	ProjectSettings                string            `json:"ProjectSettings"`
	SRU                            int               `json:"SRU"`
	TransactionInformation         string            `json:"TransactionInformation"`
	TransactionInformationSettings string            `json:"TransactionInformationSettings"`
	VATCode                        string            `json:"VATCode"`
	OpeningQuantities              []OpeningQuantity `json:"OpeningQuantities"`
}

type CreateAccountResp struct {
	Account CreateAccountInnerResp `json:"Account"`
}

type CreateAccountInnerResp struct {
	Url                            string            `json:"@url"`
	Active                         bool              `json:"Active"`
	BalanceBroughtForward          int               `json:"BalanceBroughtForward"`
	CostCenter                     string            `json:"CostCenter"`
	CostCenterSettings             string            `json:"CostCenterSettings"`
	Description                    string            `json:"Description"`
	Number                         int               `json:"Number"`
	Project                        string            `json:"Project"`
	ProjectSettings                string            `json:"ProjectSettings"`
	SRU                            int               `json:"SRU"`
	Year                           int               `json:"Year"`
	VATCode                        string            `json:"VATCode"`
	BalanceCarriedForward          int               `json:"BalanceCarriedForward"`
	TransactionInformation         string            `json:"TransactionInformation"`
	TransactionInformationSettings string            `json:"TransactionInformationSettings"`
	QuantitySettings               string            `json:"QuantitySettings"`
	QuantityUnit                   string            `json:"QuantityUnit"`
	OpeningQuantities              []OpeningQuantity `json:"OpeningQuantities"`
}
