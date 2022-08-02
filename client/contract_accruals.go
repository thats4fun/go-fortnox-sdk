package client

import (
	"context"
	"fmt"
)

const (
	contractAccrualsURI = "contractaccruals"
)

// GetAllContractAccruals does _GET https://api.fortnox.se/3/contractaccruals/
func (c *Client) GetAllContractAccruals(ctx context.Context) ([]ShortContractAccrual, error) {
	resp := &GetAllContractAccrualsResp{}

	err := c._GET(ctx, contractAccrualsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.ContractAccruals, nil
}

// CreateContractAccrual does _POST https://api.fortnox.se/3/contractaccruals/
//
// req - contract accrual to create
func (c *Client) CreateContractAccrual(ctx context.Context, fca *FullContractAccrual) (*FullContractAccrual, error) {

	req := &CreateContractAccrualReq{ContractAccrual: *fca}
	resp := &CreateContractAccrualResp{}

	err := c._POST(ctx, contractAccrualsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractAccrual, nil
}

// GetContractAccrual does _GET https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
func (c *Client) GetContractAccrual(ctx context.Context, documentNumber int) (*FullContractAccrual, error) {
	resp := &GetContractAccrualResp{}

	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractAccrual, nil
}

// UpdateContractAccrual does _PUT https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
//
// req - contract accruals to update
func (c *Client) UpdateContractAccrual(
	ctx context.Context,
	documentNumber int,
	fca *FullContractAccrual) (*FullContractAccrual, error) {

	req := &UpdateContractAccrualReq{ContractAccrual: *fca}
	resp := &UpdateContractAccrualResp{}

	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractAccrual, nil
}

// RemoveContractAccrual does _DELETE https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
func (c *Client) RemoveContractAccrual(ctx context.Context, documentNumber int) error {
	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)
	return c._DELETE(ctx, uri)
}

type FullContractAccrual struct {
	Url            string       `json:"@url,omitempty"`
	AccrualAccount int          `json:"AccrualAccount,omitempty"`
	CostAccount    int          `json:"CostAccount,omitempty"`
	Description    string       `json:"Description,omitempty"`
	AccrualRows    []AccrualRow `json:"AccrualRows,omitempty"`
	DocumentNumber int          `json:"DocumentNumber,omitempty"`
	Period         string       `json:"Period,omitempty"`
	Times          int          `json:"Times,omitempty"`
	Total          int          `json:"Total,omitempty"`
	VATIncluded    bool         `json:"VATIncluded,omitempty"`
}

type AccrualRow struct {
	Account                int    `json:"Account,omitempty"`
	CostCenter             string `json:"CostCenter,omitempty"`
	Credit                 int    `json:"Credit,omitempty"`
	Debit                  int    `json:"Debit,omitempty"`
	Project                string `json:"Project,omitempty"`
	TransactionInformation string `json:"TransactionInformation,omitempty"`
}

type ShortContractAccrual struct {
	Url            string `json:"@url,omitempty"`
	Description    string `json:"Description,omitempty"`
	DocumentNumber int    `json:"DocumentNumber,omitempty"`
	Period         string `json:"Period,omitempty"`
}

type GetAllContractAccrualsResp struct {
	ContractAccruals []ShortContractAccrual `json:"ContractAccruals"`
}

type CreateContractAccrualReq struct {
	ContractAccrual FullContractAccrual `json:"ContractAccrual"`
}

type CreateContractAccrualResp struct {
	ContractAccrual FullContractAccrual `json:"ContractAccrual"`
}

type GetContractAccrualResp struct {
	ContractAccrual FullContractAccrual `json:"ContractAccrual"`
}

type UpdateContractAccrualReq struct {
	ContractAccrual FullContractAccrual `json:"ContractAccrual"`
}

type UpdateContractAccrualResp struct {
	ContractAccrual FullContractAccrual `json:"ContractAccrual"`
}
