package client

import (
	"context"
	"fmt"
)

const (
	contractTemplatesURI = "contracttemplates"
)

// GetAllContractTemplates does _GET https://api.fortnox.se/3/contracttemplates/
func (c *Client) GetAllContractTemplates(ctx context.Context) (*GetAllContractTemplatesResp, error) {
	resp := &GetAllContractTemplatesResp{}

	err := c._GET(ctx, contractTemplatesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateContractTemplate does _POST https://api.fortnox.se/3/contracttemplates/
//
// req - contract template to create
func (c *Client) CreateContractTemplate(
	ctx context.Context,
	req *CreateContractTemplateReq) (*CreateContractTemplateResp, error) {

	resp := &CreateContractTemplateResp{}

	err := c._POST(ctx, contractTemplatesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetContractTemplate does _GET https://api.fortnox.se/3/contracttemplates/{DocumentNumber}
//
// templateNumber - identifies the contract accrual
func (c *Client) GetContractTemplate(ctx context.Context, templateNumber int) (*GetContractTemplateResp, error) {
	resp := &GetContractTemplateResp{}

	uri := fmt.Sprintf("%s/%d", contractTemplatesURI, templateNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateContractTemplate does _PUT https://api.fortnox.se/3/contracttemplates/{DocumentNumber}
//
// templateNumber - identifies the contract accrual
//
// req - contract template to update
func (c *Client) UpdateContractTemplate(
	ctx context.Context,
	templateNumber int,
	req *UpdateContractTemplateReq) (*UpdateContractTemplateResp, error) {

	resp := &UpdateContractTemplateResp{}

	uri := fmt.Sprintf("%s/%d", contractTemplatesURI, templateNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllContractTemplatesResp struct {
	ContractTemplates []struct {
		Url                  string `json:"@url"`
		ContractLength       int    `json:"ContractLength"`
		ContractTemplate     int    `json:"ContractTemplate"`
		ContractTemplateName string `json:"ContractTemplateName"`
		InvoiceInterval      int    `json:"InvoiceInterval"`
	} `json:"ContractTemplates"`
}

type CreateContractTemplateReq struct {
	ContractTemplate struct {
		Url               string `json:"@url"`
		AdministrationFee int    `json:"AdministrationFee"`
		ContractLength    int    `json:"ContractLength"`
		Freight           int    `json:"Freight"`
		InvoiceInterval   int    `json:"InvoiceInterval"`
		InvoiceRows       []struct {
			AccountNumber     int    `json:"AccountNumber"`
			ArticleNumber     string `json:"ArticleNumber"`
			CostCenter        string `json:"CostCenter"`
			DeliveredQuantity string `json:"DeliveredQuantity"`
			Description       string `json:"Description"`
			Discount          int    `json:"Discount"`
			DiscountType      string `json:"DiscountType"`
			Price             int    `json:"Price"`
			Project           string `json:"Project"`
			Unit              string `json:"Unit"`
		} `json:"InvoiceRows"`
		Continuous      bool   `json:"Continuous"`
		OurReference    string `json:"OurReference"`
		PrintTemplate   string `json:"PrintTemplate"`
		Remarks         string `json:"Remarks"`
		TemplateName    string `json:"TemplateName"`
		TemplateNumber  int    `json:"TemplateNumber"`
		TermsOfDelivery string `json:"TermsOfDelivery"`
		TermsOfPayment  string `json:"TermsOfPayment"`
		WayOfDelivery   string `json:"WayOfDelivery"`
	} `json:"ContractTemplate"`
}

type CreateContractTemplateResp struct {
	ContractTemplate struct {
		Url               string `json:"@url"`
		AdministrationFee int    `json:"AdministrationFee"`
		ContractLength    int    `json:"ContractLength"`
		Freight           int    `json:"Freight"`
		InvoiceInterval   int    `json:"InvoiceInterval"`
		InvoiceRows       []struct {
			AccountNumber     int    `json:"AccountNumber"`
			ArticleNumber     string `json:"ArticleNumber"`
			CostCenter        string `json:"CostCenter"`
			DeliveredQuantity string `json:"DeliveredQuantity"`
			Description       string `json:"Description"`
			Discount          int    `json:"Discount"`
			DiscountType      string `json:"DiscountType"`
			Price             int    `json:"Price"`
			Project           string `json:"Project"`
			Unit              string `json:"Unit"`
		} `json:"InvoiceRows"`
		Continuous      bool   `json:"Continuous"`
		OurReference    string `json:"OurReference"`
		PrintTemplate   string `json:"PrintTemplate"`
		Remarks         string `json:"Remarks"`
		TemplateName    string `json:"TemplateName"`
		TemplateNumber  int    `json:"TemplateNumber"`
		TermsOfDelivery string `json:"TermsOfDelivery"`
		TermsOfPayment  string `json:"TermsOfPayment"`
		WayOfDelivery   string `json:"WayOfDelivery"`
	} `json:"ContractTemplate"`
}

type GetContractTemplateResp struct {
	ContractTemplate struct {
		Url               string `json:"@url"`
		AdministrationFee int    `json:"AdministrationFee"`
		ContractLength    int    `json:"ContractLength"`
		Freight           int    `json:"Freight"`
		InvoiceInterval   int    `json:"InvoiceInterval"`
		InvoiceRows       []struct {
			AccountNumber     int    `json:"AccountNumber"`
			ArticleNumber     string `json:"ArticleNumber"`
			CostCenter        string `json:"CostCenter"`
			DeliveredQuantity string `json:"DeliveredQuantity"`
			Description       string `json:"Description"`
			Discount          int    `json:"Discount"`
			DiscountType      string `json:"DiscountType"`
			Price             int    `json:"Price"`
			Project           string `json:"Project"`
			Unit              string `json:"Unit"`
		} `json:"InvoiceRows"`
		Continuous      bool   `json:"Continuous"`
		OurReference    string `json:"OurReference"`
		PrintTemplate   string `json:"PrintTemplate"`
		Remarks         string `json:"Remarks"`
		TemplateName    string `json:"TemplateName"`
		TemplateNumber  int    `json:"TemplateNumber"`
		TermsOfDelivery string `json:"TermsOfDelivery"`
		TermsOfPayment  string `json:"TermsOfPayment"`
		WayOfDelivery   string `json:"WayOfDelivery"`
	} `json:"ContractTemplate"`
}

type UpdateContractTemplateReq struct {
	ContractTemplate struct {
		Url               string `json:"@url"`
		AdministrationFee int    `json:"AdministrationFee"`
		ContractLength    int    `json:"ContractLength"`
		Freight           int    `json:"Freight"`
		InvoiceInterval   int    `json:"InvoiceInterval"`
		InvoiceRows       []struct {
			AccountNumber     int    `json:"AccountNumber"`
			ArticleNumber     string `json:"ArticleNumber"`
			CostCenter        string `json:"CostCenter"`
			DeliveredQuantity string `json:"DeliveredQuantity"`
			Description       string `json:"Description"`
			Discount          int    `json:"Discount"`
			DiscountType      string `json:"DiscountType"`
			Price             int    `json:"Price"`
			Project           string `json:"Project"`
			Unit              string `json:"Unit"`
		} `json:"InvoiceRows"`
		Continuous      bool   `json:"Continuous"`
		OurReference    string `json:"OurReference"`
		PrintTemplate   string `json:"PrintTemplate"`
		Remarks         string `json:"Remarks"`
		TemplateName    string `json:"TemplateName"`
		TemplateNumber  int    `json:"TemplateNumber"`
		TermsOfDelivery string `json:"TermsOfDelivery"`
		TermsOfPayment  string `json:"TermsOfPayment"`
		WayOfDelivery   string `json:"WayOfDelivery"`
	} `json:"ContractTemplate"`
}

type UpdateContractTemplateResp struct {
	ContractTemplate struct {
		Url               string `json:"@url"`
		AdministrationFee int    `json:"AdministrationFee"`
		ContractLength    int    `json:"ContractLength"`
		Freight           int    `json:"Freight"`
		InvoiceInterval   int    `json:"InvoiceInterval"`
		InvoiceRows       []struct {
			AccountNumber     int    `json:"AccountNumber"`
			ArticleNumber     string `json:"ArticleNumber"`
			CostCenter        string `json:"CostCenter"`
			DeliveredQuantity string `json:"DeliveredQuantity"`
			Description       string `json:"Description"`
			Discount          int    `json:"Discount"`
			DiscountType      string `json:"DiscountType"`
			Price             int    `json:"Price"`
			Project           string `json:"Project"`
			Unit              string `json:"Unit"`
		} `json:"InvoiceRows"`
		Continuous      bool   `json:"Continuous"`
		OurReference    string `json:"OurReference"`
		PrintTemplate   string `json:"PrintTemplate"`
		Remarks         string `json:"Remarks"`
		TemplateName    string `json:"TemplateName"`
		TemplateNumber  int    `json:"TemplateNumber"`
		TermsOfDelivery string `json:"TermsOfDelivery"`
		TermsOfPayment  string `json:"TermsOfPayment"`
		WayOfDelivery   string `json:"WayOfDelivery"`
	} `json:"ContractTemplate"`
}
