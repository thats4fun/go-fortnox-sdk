package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoiceAccrualsURI = "supplierinvoiceaccruals"
)

func (c *Client) GetAllSupplierInvoiceAccruals(ctx context.Context) (*GetAllSupplierInvoiceAccrualsResp, error) {
	resp := &GetAllSupplierInvoiceAccrualsResp{}

	err := c._GET(ctx, supplierInvoiceAccrualsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateSupplierInvoiceAccruals(
	ctx context.Context,
	req *CreateSupplierInvoiceAccrualsReq) (*CreateSupplierInvoiceAccrualsResp, error) {

	resp := &CreateSupplierInvoiceAccrualsResp{}

	err := c._POST(ctx, supplierInvoiceAccrualsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetSupplierInvoiceAccruals(
	ctx context.Context,
	supplierInvoiceNumber int) (*GetSupplierInvoiceAccrualsResp, error) {

	resp := &GetSupplierInvoiceAccrualsResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) UpdateSupplierInvoiceAccruals(
	ctx context.Context,
	supplierInvoiceNumber int,
	req *UpdateSupplierInvoiceAccrualReq) (*UpdateSupplierInvoiceAccrualResp, error) {

	resp := &UpdateSupplierInvoiceAccrualResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteSupplierInvoiceAccruals(ctx context.Context, supplierInvoiceNumber int) error {
	uri := fmt.Sprintf("%s/%d", supplierInvoiceAccrualsURI, supplierInvoiceNumber)
	return c._DELETE(ctx, uri)
}

type GetAllSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccruals []struct {
		Url                   string `json:"@url"`
		Description           string `json:"Description"`
		SupplierInvoiceNumber int    `json:"SupplierInvoiceNumber"`
		Period                string `json:"Period"`
	} `json:"SupplierInvoiceAccruals"`
}

type CreateSupplierInvoiceAccrualsReq struct {
	SupplierInvoiceAccrual struct {
		Url                        string `json:"@url"`
		AccrualAccount             int    `json:"AccrualAccount"`
		CostAccount                int    `json:"CostAccount"`
		Description                string `json:"Description"`
		EndDate                    string `json:"EndDate"`
		SupplierInvoiceNumber      int    `json:"SupplierInvoiceNumber"`
		Period                     string `json:"Period"`
		StartDate                  string `json:"StartDate"`
		Times                      int    `json:"Times"`
		Total                      int    `json:"Total"`
		VATIncluded                bool   `json:"VATIncluded"`
		SupplierInvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"SupplierInvoiceAccrualRows"`
	} `json:"SupplierInvoiceAccrual"`
}

type CreateSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccrual struct {
		Url                        string `json:"@url"`
		AccrualAccount             int    `json:"AccrualAccount"`
		CostAccount                int    `json:"CostAccount"`
		Description                string `json:"Description"`
		EndDate                    string `json:"EndDate"`
		SupplierInvoiceNumber      int    `json:"SupplierInvoiceNumber"`
		Period                     string `json:"Period"`
		StartDate                  string `json:"StartDate"`
		Times                      int    `json:"Times"`
		Total                      int    `json:"Total"`
		VATIncluded                bool   `json:"VATIncluded"`
		SupplierInvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"SupplierInvoiceAccrualRows"`
	} `json:"SupplierInvoiceAccrual"`
}

type GetSupplierInvoiceAccrualsResp struct {
	SupplierInvoiceAccrual struct {
		Url                        string `json:"@url"`
		AccrualAccount             int    `json:"AccrualAccount"`
		CostAccount                int    `json:"CostAccount"`
		Description                string `json:"Description"`
		EndDate                    string `json:"EndDate"`
		SupplierInvoiceNumber      int    `json:"SupplierInvoiceNumber"`
		Period                     string `json:"Period"`
		StartDate                  string `json:"StartDate"`
		Times                      int    `json:"Times"`
		Total                      int    `json:"Total"`
		VATIncluded                bool   `json:"VATIncluded"`
		SupplierInvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"SupplierInvoiceAccrualRows"`
	} `json:"SupplierInvoiceAccrual"`
}

type UpdateSupplierInvoiceAccrualReq struct {
	SupplierInvoiceAccrual struct {
		Url                        string `json:"@url"`
		AccrualAccount             int    `json:"AccrualAccount"`
		CostAccount                int    `json:"CostAccount"`
		Description                string `json:"Description"`
		EndDate                    string `json:"EndDate"`
		SupplierInvoiceNumber      int    `json:"SupplierInvoiceNumber"`
		Period                     string `json:"Period"`
		StartDate                  string `json:"StartDate"`
		Times                      int    `json:"Times"`
		Total                      int    `json:"Total"`
		VATIncluded                bool   `json:"VATIncluded"`
		SupplierInvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"SupplierInvoiceAccrualRows"`
	} `json:"SupplierInvoiceAccrual"`
}

type UpdateSupplierInvoiceAccrualResp struct {
	SupplierInvoiceAccrual struct {
		Url                        string `json:"@url"`
		AccrualAccount             int    `json:"AccrualAccount"`
		CostAccount                int    `json:"CostAccount"`
		Description                string `json:"Description"`
		EndDate                    string `json:"EndDate"`
		SupplierInvoiceNumber      int    `json:"SupplierInvoiceNumber"`
		Period                     string `json:"Period"`
		StartDate                  string `json:"StartDate"`
		Times                      int    `json:"Times"`
		Total                      int    `json:"Total"`
		VATIncluded                bool   `json:"VATIncluded"`
		SupplierInvoiceAccrualRows []struct {
			Account                int    `json:"Account"`
			CostCenter             string `json:"CostCenter"`
			Credit                 int    `json:"Credit"`
			Debit                  int    `json:"Debit"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
		} `json:"SupplierInvoiceAccrualRows"`
	} `json:"SupplierInvoiceAccrual"`
}
