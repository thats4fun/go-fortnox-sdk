package client

import (
	"context"
	"fmt"
)

const (
	suppliersURI = "suppliers"
)

// GetAllSuppliers does _GET https://api.fortnox.se/3/suppliers
func (c *Client) GetAllSuppliers(ctx context.Context) ([]Supplier, error) {
	resp := &GetAllSuppliersResp{}

	err := c._GET(ctx, suppliersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Suppliers, nil
}

// CreateSupplier does _POST https://api.fortnox.se/3/suppliers
//
// req - supplier to create
func (c *Client) CreateSupplier(ctx context.Context, s *Supplier) (*Supplier, error) {
	req := &CreateSupplierReq{Supplier: *s}
	resp := &CreateSupplierResp{}

	err := c._POST(ctx, suppliersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Supplier, nil
}

// GetSupplier does _GET https://api.fortnox.se/3/suppliers/{SupplierNumber}
//
// supplierNumber - identifies the supplier
func (c *Client) GetSupplier(ctx context.Context, supplierNumber string) (*Supplier, error) {
	resp := &GetSupplierResp{}

	uri := fmt.Sprintf("%s/%s", suppliersURI, supplierNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Supplier, nil
}

// UpdateSupplier does _PUT https://api.fortnox.se/3/suppliers/{SupplierNumber}
//
// supplierNumber - identifies the supplier
func (c *Client) UpdateSupplier(ctx context.Context, supplierNumber string, s *Supplier) (*Supplier, error) {
	req := &UpdateSupplierReq{Supplier: *s}
	resp := &UpdateSupplierResp{}

	uri := fmt.Sprintf("%s/%s", suppliersURI, supplierNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Supplier, nil
}

type GetAllSuppliersResp struct {
	Suppliers []Supplier `json:"Suppliers"`
}

type CreateSupplierReq struct {
	Supplier Supplier `json:"Supplier"`
}

type CreateSupplierResp struct {
	Supplier Supplier `json:"Supplier"`
}

type GetSupplierResp struct {
	Supplier Supplier `json:"Supplier"`
}

type UpdateSupplierReq struct {
	Supplier Supplier `json:"Supplier"`
}

type UpdateSupplierResp struct {
	Supplier Supplier `json:"Supplier"`
}

type Supplier struct {
	Url                 string `json:"@url,omitempty"`
	Active              bool   `json:"Active,omitempty"`
	Address1            string `json:"Address1,omitempty"`
	Address2            string `json:"Address2,omitempty"`
	Bank                string `json:"Bank,omitempty"`
	BankAccountNumber   string `json:"BankAccountNumber,omitempty"`
	BG                  string `json:"BG,omitempty"`
	BIC                 string `json:"BIC,omitempty"`
	BranchCode          string `json:"BranchCode,omitempty"`
	City                string `json:"City,omitempty"`
	ClearingNumber      string `json:"ClearingNumber,omitempty"`
	Comments            string `json:"Comments,omitempty"`
	CostCenter          string `json:"CostCenter,omitempty"`
	Country             string `json:"Country,omitempty"`
	CountryCode         string `json:"CountryCode,omitempty"`
	Currency            string `json:"Currency,omitempty"`
	DisablePaymentFile  bool   `json:"DisablePaymentFile,omitempty"`
	Email               string `json:"Email,omitempty"`
	Fax                 string `json:"Fax,omitempty"`
	IBAN                string `json:"IBAN,omitempty"`
	Name                string `json:"Name,omitempty"`
	OrganisationNumber  string `json:"OrganisationNumber,omitempty"`
	OurReference        string `json:"OurReference,omitempty"`
	OurCustomerNumber   string `json:"OurCustomerNumber,omitempty"`
	PG                  string `json:"PG,omitempty"`
	Phone1              string `json:"Phone1,omitempty"`
	Phone2              string `json:"Phone2,omitempty"`
	PreDefinedAccount   string `json:"PreDefinedAccount,omitempty"`
	Project             string `json:"Project,omitempty"`
	SupplierNumber      string `json:"SupplierNumber,omitempty"`
	TermsOfPayment      string `json:"TermsOfPayment,omitempty"`
	VATNumber           string `json:"VATNumber,omitempty"`
	VATType             string `json:"VATType,omitempty"`
	VisitingAddress     string `json:"VisitingAddress,omitempty"`
	VisitingCity        string `json:"VisitingCity,omitempty"`
	VisitingCountry     string `json:"VisitingCountry,omitempty"`
	VisitingCountryCode string `json:"VisitingCountryCode,omitempty"`
	VisitingZipCode     string `json:"VisitingZipCode,omitempty"`
	WorkPlace           string `json:"WorkPlace,omitempty"`
	WWW                 string `json:"WWW,omitempty"`
	YourReference       string `json:"YourReference,omitempty"`
	ZipCode             string `json:"ZipCode,omitempty"`
}
