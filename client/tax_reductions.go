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
func (c Client) GetAllTaxReductions(ctx context.Context, filter *GetAllTaxReductionsFilter) ([]TaxReduction, error) {
	resp := &GetAllTaxReductionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, taxReductionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.TaxReductions, nil
}

// CreateTaxReduction does _POST https://api.fortnox.se/3/taxreductions
//
// req - tax reduction to create
func (c *Client) CreateTaxReduction(ctx context.Context, tr *TaxReduction) (*TaxReduction, error) {
	req := &CreateTaxReductionReq{TaxReduction: *tr}
	resp := &CreateTaxReductionResp{}

	err := c._POST(ctx, taxReductionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TaxReduction, nil
}

// GetTaxReduction does _GET https://api.fortnox.se/3/taxreductions/{Id}
//
// id - identifies the tax reduction
func (c *Client) GetTaxReduction(ctx context.Context, id int) (*TaxReduction, error) {
	resp := &GetTaxReductionResp{}

	uri := fmt.Sprintf("%s/%d", taxReductionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TaxReduction, nil
}

// UpdateTaxReduction does _PUT https://api.fortnox.se/3/taxreductions/{Id}
//
// id - identifies the tax reduction
//
// req - tax reduction to update
func (c *Client) UpdateTaxReduction(ctx context.Context, id int, tr *TaxReduction) (*TaxReduction, error) {
	req := &UpdateTaxReductionReq{TaxReduction: *tr}
	resp := &UpdateTaxReductionResp{}

	uri := fmt.Sprintf("%s/%d", taxReductionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.TaxReduction, nil
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

type TaxReduction struct {
	Url                                    string               `json:"@url,omitempty"`
	ApprovedAmount                         int                  `json:"ApprovedAmount,omitempty"`
	AskedAmount                            int                  `json:"AskedAmount,omitempty"`
	BilledAmount                           int                  `json:"BilledAmount,omitempty"`
	CustomerName                           string               `json:"CustomerName,omitempty"`
	Id                                     int                  `json:"Id,omitempty"`
	PropertyDesignation                    string               `json:"PropertyDesignation,omitempty"`
	ReferenceDocumentType                  string               `json:"ReferenceDocumentType,omitempty"`
	ReferenceNumber                        string               `json:"ReferenceNumber,omitempty"`
	RequestSent                            bool                 `json:"RequestSent,omitempty"`
	ResidenceAssociationOrganisationNumber string               `json:"ResidenceAssociationOrganisationNumber,omitempty"`
	SocialSecurityNumber                   string               `json:"SocialSecurityNumber,omitempty"`
	TypeOfReduction                        string               `json:"TypeOfReduction,omitempty"`
	VoucherNumber                          int                  `json:"VoucherNumber,omitempty"`
	VoucherSeries                          string               `json:"VoucherSeries,omitempty"`
	VoucherYear                            int                  `json:"VoucherYear,omitempty"`
	TaxReductionAmounts                    []TaxReductionAmount `json:"TaxReductionAmounts,omitempty"`
}

type TaxReductionAmount struct {
	AskedAmount int    `json:"askedAmount,omitempty"`
	WorkType    string `json:"workType,omitempty"`
}

type GetAllTaxReductionsResp struct {
	TaxReductions []TaxReduction `json:"TaxReductions"`
}

type CreateTaxReductionReq struct {
	TaxReduction TaxReduction `json:"TaxReduction"`
}

type CreateTaxReductionResp struct {
	TaxReduction TaxReduction `json:"TaxReduction"`
}

type GetTaxReductionResp struct {
	TaxReduction TaxReduction `json:"TaxReduction"`
}

type UpdateTaxReductionReq struct {
	TaxReduction TaxReduction `json:"TaxReduction"`
}

type UpdateTaxReductionResp struct {
	TaxReduction TaxReduction `json:"TaxReduction"`
}
