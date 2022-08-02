package client

import (
	"context"
	"fmt"
)

const (
	contractTemplatesURI = "contracttemplates"
)

// GetAllContractTemplates does _GET https://api.fortnox.se/3/contracttemplates/
func (c *Client) GetAllContractTemplates(ctx context.Context) ([]ContractTemplate, error) {
	resp := &GetAllContractTemplatesResp{}

	err := c._GET(ctx, contractTemplatesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.ContractTemplates, nil
}

// CreateContractTemplate does _POST https://api.fortnox.se/3/contracttemplates/
//
// ct - contract template to create
func (c *Client) CreateContractTemplate(ctx context.Context, ct *ContractTemplate) (*ContractTemplate, error) {
	req := &CreateContractTemplateReq{ContractTemplate: *ct}
	resp := &CreateContractTemplateResp{}

	err := c._POST(ctx, contractTemplatesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractTemplate, nil
}

// GetContractTemplate does _GET https://api.fortnox.se/3/contracttemplates/{DocumentNumber}
//
// templateNumber - identifies the contract accrual
func (c *Client) GetContractTemplate(ctx context.Context, templateNumber int) (*ContractTemplate, error) {
	resp := &GetContractTemplateResp{}

	uri := fmt.Sprintf("%s/%d", contractTemplatesURI, templateNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractTemplate, nil
}

// UpdateContractTemplate does _PUT https://api.fortnox.se/3/contracttemplates/{DocumentNumber}
//
// templateNumber - identifies the contract accrual
//
// ct - contract template to update
func (c *Client) UpdateContractTemplate(
	ctx context.Context,
	templateNumber int,
	ct *ContractTemplate) (*ContractTemplate, error) {

	req := &UpdateContractTemplateReq{ContractTemplate: *ct}
	resp := &UpdateContractTemplateResp{}

	uri := fmt.Sprintf("%s/%d", contractTemplatesURI, templateNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ContractTemplate, nil
}

type GetAllContractTemplatesResp struct {
	ContractTemplates []ContractTemplate `json:"ContractTemplates"`
}

type CreateContractTemplateReq struct {
	ContractTemplate ContractTemplate `json:"ContractTemplate"`
}

type CreateContractTemplateResp struct {
	ContractTemplate ContractTemplate `json:"ContractTemplate"`
}

type GetContractTemplateResp struct {
	ContractTemplate ContractTemplate `json:"ContractTemplate"`
}

type UpdateContractTemplateReq struct {
	ContractTemplate ContractTemplate `json:"ContractTemplate"`
}

type UpdateContractTemplateResp struct {
	ContractTemplate ContractTemplate `json:"ContractTemplate"`
}

type ContractTemplate struct {
	Url               string       `json:"@url,omitempty"`
	AdministrationFee int          `json:"AdministrationFee,omitempty"`
	ContractLength    int          `json:"ContractLength,omitempty"`
	Freight           int          `json:"Freight,omitempty"`
	InvoiceInterval   int          `json:"InvoiceInterval,omitempty"`
	InvoiceRows       []InvoiceRow `json:"InvoiceRows,omitempty"`
	Continuous        bool         `json:"Continuous,omitempty"`
	OurReference      string       `json:"OurReference,omitempty"`
	PrintTemplate     string       `json:"PrintTemplate,omitempty"`
	Remarks           string       `json:"Remarks,omitempty"`
	TemplateName      string       `json:"TemplateName,omitempty"`
	TemplateNumber    int          `json:"TemplateNumber,omitempty"`
	TermsOfDelivery   string       `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment    string       `json:"TermsOfPayment,omitempty"`
	WayOfDelivery     string       `json:"WayOfDelivery,omitempty"`
}
