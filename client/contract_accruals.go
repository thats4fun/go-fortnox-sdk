package client

import (
	"context"
	"fmt"
)

const (
	contractAccrualsURI = "contractaccruals"
)

// GetAllContractAccruals does _GET https://api.fortnox.se/3/contractaccruals/
func (c *Client) GetAllContractAccruals(ctx context.Context) (*GetAllContractAccrualsResp, error) {
	resp := &GetAllContractAccrualsResp{}

	err := c._GET(ctx, contractAccrualsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateContractAccrual does _POST https://api.fortnox.se/3/contractaccruals/
//
// req - contract accrual to create
func (c *Client) CreateContractAccrual(
	ctx context.Context,
	req *CreateContractAccrualReq) (*CreateContractAccrualResp, error) {

	resp := &CreateContractAccrualResp{}

	err := c._POST(ctx, contractAccrualsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetContractAccrual does _GET https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
func (c *Client) GetContractAccrual(ctx context.Context, documentNumber int) (*GetContractAccrualResp, error) {
	resp := &GetContractAccrualResp{}

	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateContractAccrual does _PUT https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
//
// req - contract accruals to update
func (c *Client) UpdateContractAccrual(
	ctx context.Context,
	documentNumber int,
	req *UpdateContractAccrualReq) (*UpdateContractAccrualResp, error) {

	resp := &UpdateContractAccrualResp{}

	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveContractAccrual does DELET https://api.fortnox.se/3/contractaccruals/{DocumentNumber}
//
// documentNumber - identifies the contract accrual
func (c *Client) RemoveContractAccrual(ctx context.Context, documentNumber int) error {
	uri := fmt.Sprintf("%s/%d", contractAccrualsURI, documentNumber)
	return c._DELETE(ctx, uri)
}

type GetAllContractAccrualsResp struct {
	ContractAccruals []struct {
		Url            string `json:"@url"`
		Description    string `json:"Description"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
	} `json:"ContractAccruals"`
}

type CreateContractAccrualReq struct {
	ContractAccrual struct {
		Url            string `json:"@url"`
		AccrualAccount int    `json:"AccrualAccount"`
		CostAccount    int    `json:"CostAccount"`
		Description    string `json:"Description"`
		AccrualRows    []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"AccrualRows"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"ContractAccrual"`
}

type CreateContractAccrualResp struct {
	ContractAccrual struct {
		Url            string `json:"@url"`
		AccrualAccount int    `json:"AccrualAccount"`
		CostAccount    int    `json:"CostAccount"`
		Description    string `json:"Description"`
		AccrualRows    []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"AccrualRows"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"ContractAccrual"`
}

type GetContractAccrualResp struct {
	ContractAccrual struct {
		Url            string `json:"@url"`
		AccrualAccount int    `json:"AccrualAccount"`
		CostAccount    int    `json:"CostAccount"`
		Description    string `json:"Description"`
		AccrualRows    []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"AccrualRows"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"ContractAccrual"`
}

type UpdateContractAccrualReq struct {
	ContractAccrual struct {
		Url            string `json:"@url"`
		AccrualAccount int    `json:"AccrualAccount"`
		CostAccount    int    `json:"CostAccount"`
		Description    string `json:"Description"`
		AccrualRows    []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"AccrualRows"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"ContractAccrual"`
}

type UpdateContractAccrualResp struct {
	ContractAccrual struct {
		Url            string `json:"@url"`
		AccrualAccount int    `json:"AccrualAccount"`
		CostAccount    int    `json:"CostAccount"`
		Description    string `json:"Description"`
		AccrualRows    []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"AccrualRows"`
		DocumentNumber int    `json:"DocumentNumber"`
		Period         string `json:"Period"`
		Times          int    `json:"Times"`
		Total          int    `json:"Total"`
		VATIncluded    bool   `json:"VATIncluded"`
	} `json:"ContractAccrual"`
}
