package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	taxReductionsURI = "taxreductions"
)

const (
	taxReductionFilterName = "filter"
)

// GetAllTaxReductions does _GET https://api.fortnox.se/3/taxreductions
//
// Enum: {"invoices" "orders" "offers"}, possibility to filter tax reductions
func (c Client) GetAllTaxReductions(
	ctx context.Context,
	filter *GetAllTaxReductionsFilter) (*GetAllTaxReductionsResp, error) {

	resp := &GetAllTaxReductionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, taxReductionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateTaxReduction does _POST https://api.fortnox.se/3/taxreductions
//
// req - tax reduction to create
func (c *Client) CreateTaxReduction(ctx context.Context, req *CreateTaxReductionReq) (*CreateTaxReductionResp, error) {
	resp := &CreateTaxReductionResp{}

	err := c._POST(ctx, taxReductionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTaxReduction does _GET https://api.fortnox.se/3/taxreductions/{Id}
//
// id - identifies the tax reduction
func (c *Client) GetTaxReduction(ctx context.Context, id int) (*GetTaxReductionResp, error) {
	resp := &GetTaxReductionResp{}

	uri := fmt.Sprintf("%s/%d", taxReductionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateTaxReduction does _PUT https://api.fortnox.se/3/taxreductions/{Id}
//
// id - identifies the tax reduction
//
// req - tax reduction to update
func (c *Client) UpdateTaxReduction(
	ctx context.Context,
	id int,
	req *UpdateTaxReductionReq) (*UpdateTaxReductionResp, error) {

	resp := &UpdateTaxReductionResp{}

	uri := fmt.Sprintf("%s/%d", taxReductionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveTaxReduction does _DELETE https://api.fortnox.se/3/taxreductions/{Id}
//
// id - identifies the tax reduction
func (c Client) RemoveTaxReduction(ctx context.Context, id int) error {
	uri := fmt.Sprintf("%s/%d", taxReductionsURI, id)
	return c._DELETE(ctx, uri)
}

type GetAllTaxReductionsFilter string

const (
	InvoicesTaxReductionFilter GetAllTaxReductionsFilter = "invoices"
	OrdersTaxReductionFilter   GetAllTaxReductionsFilter = "orders"
	OffersTaxReductionFilter   GetAllTaxReductionsFilter = "invoices"
)

func (f *GetAllTaxReductionsFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	fStr := string(*f)

	if strings.TrimSpace(fStr) != "" {
		params[taxReductionFilterName] = []string{fStr}
	}

	return params
}

type GetAllTaxReductionsResp struct {
	TaxReductions []struct {
		Url                   string `json:"@url"`
		ApprovedAmount        int    `json:"ApprovedAmount"`
		CustomerName          string `json:"CustomerName"`
		Id                    int    `json:"Id"`
		ReferenceDocumentType string `json:"ReferenceDocumentType"`
		ReferenceNumber       int    `json:"ReferenceNumber"`
		SocialSecurityNumber  string `json:"SocialSecurityNumber"`
	} `json:"TaxReductions"`
}

type CreateTaxReductionReq struct {
	TaxReduction struct {
		Url                                    string `json:"@url"`
		ApprovedAmount                         int    `json:"ApprovedAmount"`
		AskedAmount                            int    `json:"AskedAmount"`
		BilledAmount                           int    `json:"BilledAmount"`
		CustomerName                           string `json:"CustomerName"`
		Id                                     int    `json:"Id"`
		PropertyDesignation                    string `json:"PropertyDesignation"`
		ReferenceDocumentType                  string `json:"ReferenceDocumentType"`
		ReferenceNumber                        string `json:"ReferenceNumber"`
		RequestSent                            bool   `json:"RequestSent"`
		ResidenceAssociationOrganisationNumber string `json:"ResidenceAssociationOrganisationNumber"`
		SocialSecurityNumber                   string `json:"SocialSecurityNumber"`
		TypeOfReduction                        string `json:"TypeOfReduction"`
		VoucherNumber                          int    `json:"VoucherNumber"`
		VoucherSeries                          string `json:"VoucherSeries"`
		VoucherYear                            int    `json:"VoucherYear"`
		TaxReductionAmounts                    []struct {
			AskedAmount int    `json:"askedAmount"`
			WorkType    string `json:"workType"`
		} `json:"TaxReductionAmounts"`
	} `json:"TaxReduction"`
}

type CreateTaxReductionResp struct {
	TaxReduction struct {
		Url                                    string `json:"@url"`
		ApprovedAmount                         int    `json:"ApprovedAmount"`
		AskedAmount                            int    `json:"AskedAmount"`
		BilledAmount                           int    `json:"BilledAmount"`
		CustomerName                           string `json:"CustomerName"`
		Id                                     int    `json:"Id"`
		PropertyDesignation                    string `json:"PropertyDesignation"`
		ReferenceDocumentType                  string `json:"ReferenceDocumentType"`
		ReferenceNumber                        string `json:"ReferenceNumber"`
		RequestSent                            bool   `json:"RequestSent"`
		ResidenceAssociationOrganisationNumber string `json:"ResidenceAssociationOrganisationNumber"`
		SocialSecurityNumber                   string `json:"SocialSecurityNumber"`
		TypeOfReduction                        string `json:"TypeOfReduction"`
		VoucherNumber                          int    `json:"VoucherNumber"`
		VoucherSeries                          string `json:"VoucherSeries"`
		VoucherYear                            int    `json:"VoucherYear"`
		TaxReductionAmounts                    []struct {
			AskedAmount int    `json:"askedAmount"`
			WorkType    string `json:"workType"`
		} `json:"TaxReductionAmounts"`
	} `json:"TaxReduction"`
}

type GetTaxReductionResp struct {
	TaxReduction struct {
		Url                                    string `json:"@url"`
		ApprovedAmount                         int    `json:"ApprovedAmount"`
		AskedAmount                            int    `json:"AskedAmount"`
		BilledAmount                           int    `json:"BilledAmount"`
		CustomerName                           string `json:"CustomerName"`
		Id                                     int    `json:"Id"`
		PropertyDesignation                    string `json:"PropertyDesignation"`
		ReferenceDocumentType                  string `json:"ReferenceDocumentType"`
		ReferenceNumber                        string `json:"ReferenceNumber"`
		RequestSent                            bool   `json:"RequestSent"`
		ResidenceAssociationOrganisationNumber string `json:"ResidenceAssociationOrganisationNumber"`
		SocialSecurityNumber                   string `json:"SocialSecurityNumber"`
		TypeOfReduction                        string `json:"TypeOfReduction"`
		VoucherNumber                          int    `json:"VoucherNumber"`
		VoucherSeries                          string `json:"VoucherSeries"`
		VoucherYear                            int    `json:"VoucherYear"`
		TaxReductionAmounts                    []struct {
			AskedAmount int    `json:"askedAmount"`
			WorkType    string `json:"workType"`
		} `json:"TaxReductionAmounts"`
	} `json:"TaxReduction"`
}

type UpdateTaxReductionReq struct {
	TaxReduction struct {
		Url                                    string `json:"@url"`
		ApprovedAmount                         int    `json:"ApprovedAmount"`
		AskedAmount                            int    `json:"AskedAmount"`
		BilledAmount                           int    `json:"BilledAmount"`
		CustomerName                           string `json:"CustomerName"`
		Id                                     int    `json:"Id"`
		PropertyDesignation                    string `json:"PropertyDesignation"`
		ReferenceDocumentType                  string `json:"ReferenceDocumentType"`
		ReferenceNumber                        string `json:"ReferenceNumber"`
		RequestSent                            bool   `json:"RequestSent"`
		ResidenceAssociationOrganisationNumber string `json:"ResidenceAssociationOrganisationNumber"`
		SocialSecurityNumber                   string `json:"SocialSecurityNumber"`
		TypeOfReduction                        string `json:"TypeOfReduction"`
		VoucherNumber                          int    `json:"VoucherNumber"`
		VoucherSeries                          string `json:"VoucherSeries"`
		VoucherYear                            int    `json:"VoucherYear"`
		TaxReductionAmounts                    []struct {
			AskedAmount int    `json:"askedAmount"`
			WorkType    string `json:"workType"`
		} `json:"TaxReductionAmounts"`
	} `json:"TaxReduction"`
}

type UpdateTaxReductionResp struct {
	TaxReduction struct {
		Url                                    string `json:"@url"`
		ApprovedAmount                         int    `json:"ApprovedAmount"`
		AskedAmount                            int    `json:"AskedAmount"`
		BilledAmount                           int    `json:"BilledAmount"`
		CustomerName                           string `json:"CustomerName"`
		Id                                     int    `json:"Id"`
		PropertyDesignation                    string `json:"PropertyDesignation"`
		ReferenceDocumentType                  string `json:"ReferenceDocumentType"`
		ReferenceNumber                        string `json:"ReferenceNumber"`
		RequestSent                            bool   `json:"RequestSent"`
		ResidenceAssociationOrganisationNumber string `json:"ResidenceAssociationOrganisationNumber"`
		SocialSecurityNumber                   string `json:"SocialSecurityNumber"`
		TypeOfReduction                        string `json:"TypeOfReduction"`
		VoucherNumber                          int    `json:"VoucherNumber"`
		VoucherSeries                          string `json:"VoucherSeries"`
		VoucherYear                            int    `json:"VoucherYear"`
		TaxReductionAmounts                    []struct {
			AskedAmount int    `json:"askedAmount"`
			WorkType    string `json:"workType"`
		} `json:"TaxReductionAmounts"`
	} `json:"TaxReduction"`
}
