package client

import (
	"context"
	"fmt"
)

const (
	suppliersURI = "suppliers"
)

// GetAllSuppliers does _GET
func (c *Client) GetAllSuppliers(ctx context.Context) (*GetAllSuppliersResp, error) {
	resp := &GetAllSuppliersResp{}

	err := c._GET(ctx, suppliersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSupplier does _POST
//
// req - to create
func (c *Client) CreateSupplier(ctx context.Context, req *CreateSupplierReq) (*CreateSupplierResp, error) {
	resp := &CreateSupplierResp{}

	err := c._POST(ctx, suppliersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSupplier does _GET
//
// supplierNumber - identifies the supplier
func (c *Client) GetSupplier(ctx context.Context, supplierNumber string) (*GetSupplierResp, error) {
	resp := &GetSupplierResp{}

	uri := fmt.Sprintf("%s/%s", suppliersURI, supplierNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSupplier does _PUT
//
// supplierNumber - identifies the supplier
func (c *Client) UpdateSupplier(
	ctx context.Context,
	supplierNumber string,
	req *UpdateSupplierReq) (*UpdateSupplierResp, error) {

	resp := &UpdateSupplierResp{}

	uri := fmt.Sprintf("%s/%s", suppliersURI, supplierNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllSuppliersResp struct {
	Suppliers []struct {
		Url                string `json:"@url"`
		Active             bool   `json:"Active"`
		Address1           string `json:"Address1"`
		Address2           string `json:"Address2"`
		BankAccountNumber  string `json:"BankAccountNumber"`
		BG                 string `json:"BG"`
		BIC                string `json:"BIC"`
		City               string `json:"City"`
		CostCenter         string `json:"CostCenter"`
		CountryCode        string `json:"CountryCode"`
		Currency           string `json:"Currency"`
		DisablePaymentFile bool   `json:"DisablePaymentFile"`
		Email              string `json:"Email"`
		IBAN               string `json:"IBAN"`
		Name               string `json:"Name"`
		OrganisationNumber string `json:"OrganisationNumber"`
		PG                 string `json:"PG"`
		Phone              string `json:"Phone"`
		PreDefinedAccount  string `json:"PreDefinedAccount"`
		Project            string `json:"Project"`
		SupplierNumber     string `json:"SupplierNumber"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		ZipCode            string `json:"ZipCode"`
	} `json:"Suppliers"`
}

type CreateSupplierReq struct {
	Supplier struct {
		Url                 string `json:"@url"`
		Active              bool   `json:"Active"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Bank                string `json:"Bank"`
		BankAccountNumber   string `json:"BankAccountNumber"`
		BG                  string `json:"BG"`
		BIC                 string `json:"BIC"`
		BranchCode          string `json:"BranchCode"`
		City                string `json:"City"`
		ClearingNumber      string `json:"ClearingNumber"`
		Comments            string `json:"Comments"`
		CostCenter          string `json:"CostCenter"`
		Country             string `json:"Country"`
		CountryCode         string `json:"CountryCode"`
		Currency            string `json:"Currency"`
		DisablePaymentFile  bool   `json:"DisablePaymentFile"`
		Email               string `json:"Email"`
		Fax                 string `json:"Fax"`
		IBAN                string `json:"IBAN"`
		Name                string `json:"Name"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		OurCustomerNumber   string `json:"OurCustomerNumber"`
		PG                  string `json:"PG"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PreDefinedAccount   string `json:"PreDefinedAccount"`
		Project             string `json:"Project"`
		SupplierNumber      string `json:"SupplierNumber"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		VATNumber           string `json:"VATNumber"`
		VATType             string `json:"VATType"`
		VisitingAddress     string `json:"VisitingAddress"`
		VisitingCity        string `json:"VisitingCity"`
		VisitingCountry     string `json:"VisitingCountry"`
		VisitingCountryCode string `json:"VisitingCountryCode"`
		VisitingZipCode     string `json:"VisitingZipCode"`
		WorkPlace           string `json:"WorkPlace"`
		WWW                 string `json:"WWW"`
		YourReference       string `json:"YourReference"`
		ZipCode             string `json:"ZipCode"`
	} `json:"Supplier"`
}

type CreateSupplierResp struct {
	Supplier struct {
		Url                 string `json:"@url"`
		Active              bool   `json:"Active"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Bank                string `json:"Bank"`
		BankAccountNumber   string `json:"BankAccountNumber"`
		BG                  string `json:"BG"`
		BIC                 string `json:"BIC"`
		BranchCode          string `json:"BranchCode"`
		City                string `json:"City"`
		ClearingNumber      string `json:"ClearingNumber"`
		Comments            string `json:"Comments"`
		CostCenter          string `json:"CostCenter"`
		Country             string `json:"Country"`
		CountryCode         string `json:"CountryCode"`
		Currency            string `json:"Currency"`
		DisablePaymentFile  bool   `json:"DisablePaymentFile"`
		Email               string `json:"Email"`
		Fax                 string `json:"Fax"`
		IBAN                string `json:"IBAN"`
		Name                string `json:"Name"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		OurCustomerNumber   string `json:"OurCustomerNumber"`
		PG                  string `json:"PG"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PreDefinedAccount   string `json:"PreDefinedAccount"`
		Project             string `json:"Project"`
		SupplierNumber      string `json:"SupplierNumber"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		VATNumber           string `json:"VATNumber"`
		VATType             string `json:"VATType"`
		VisitingAddress     string `json:"VisitingAddress"`
		VisitingCity        string `json:"VisitingCity"`
		VisitingCountry     string `json:"VisitingCountry"`
		VisitingCountryCode string `json:"VisitingCountryCode"`
		VisitingZipCode     string `json:"VisitingZipCode"`
		WorkPlace           string `json:"WorkPlace"`
		WWW                 string `json:"WWW"`
		YourReference       string `json:"YourReference"`
		ZipCode             string `json:"ZipCode"`
	} `json:"Supplier"`
}

type GetSupplierResp struct {
	Supplier struct {
		Url                 string `json:"@url"`
		Active              bool   `json:"Active"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Bank                string `json:"Bank"`
		BankAccountNumber   string `json:"BankAccountNumber"`
		BG                  string `json:"BG"`
		BIC                 string `json:"BIC"`
		BranchCode          string `json:"BranchCode"`
		City                string `json:"City"`
		ClearingNumber      string `json:"ClearingNumber"`
		Comments            string `json:"Comments"`
		CostCenter          string `json:"CostCenter"`
		Country             string `json:"Country"`
		CountryCode         string `json:"CountryCode"`
		Currency            string `json:"Currency"`
		DisablePaymentFile  bool   `json:"DisablePaymentFile"`
		Email               string `json:"Email"`
		Fax                 string `json:"Fax"`
		IBAN                string `json:"IBAN"`
		Name                string `json:"Name"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		OurCustomerNumber   string `json:"OurCustomerNumber"`
		PG                  string `json:"PG"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PreDefinedAccount   string `json:"PreDefinedAccount"`
		Project             string `json:"Project"`
		SupplierNumber      string `json:"SupplierNumber"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		VATNumber           string `json:"VATNumber"`
		VATType             string `json:"VATType"`
		VisitingAddress     string `json:"VisitingAddress"`
		VisitingCity        string `json:"VisitingCity"`
		VisitingCountry     string `json:"VisitingCountry"`
		VisitingCountryCode string `json:"VisitingCountryCode"`
		VisitingZipCode     string `json:"VisitingZipCode"`
		WorkPlace           string `json:"WorkPlace"`
		WWW                 string `json:"WWW"`
		YourReference       string `json:"YourReference"`
		ZipCode             string `json:"ZipCode"`
	} `json:"Supplier"`
}

type UpdateSupplierReq struct {
	Supplier struct {
		Url                 string `json:"@url"`
		Active              bool   `json:"Active"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Bank                string `json:"Bank"`
		BankAccountNumber   string `json:"BankAccountNumber"`
		BG                  string `json:"BG"`
		BIC                 string `json:"BIC"`
		BranchCode          string `json:"BranchCode"`
		City                string `json:"City"`
		ClearingNumber      string `json:"ClearingNumber"`
		Comments            string `json:"Comments"`
		CostCenter          string `json:"CostCenter"`
		Country             string `json:"Country"`
		CountryCode         string `json:"CountryCode"`
		Currency            string `json:"Currency"`
		DisablePaymentFile  bool   `json:"DisablePaymentFile"`
		Email               string `json:"Email"`
		Fax                 string `json:"Fax"`
		IBAN                string `json:"IBAN"`
		Name                string `json:"Name"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		OurCustomerNumber   string `json:"OurCustomerNumber"`
		PG                  string `json:"PG"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PreDefinedAccount   string `json:"PreDefinedAccount"`
		Project             string `json:"Project"`
		SupplierNumber      string `json:"SupplierNumber"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		VATNumber           string `json:"VATNumber"`
		VATType             string `json:"VATType"`
		VisitingAddress     string `json:"VisitingAddress"`
		VisitingCity        string `json:"VisitingCity"`
		VisitingCountry     string `json:"VisitingCountry"`
		VisitingCountryCode string `json:"VisitingCountryCode"`
		VisitingZipCode     string `json:"VisitingZipCode"`
		WorkPlace           string `json:"WorkPlace"`
		WWW                 string `json:"WWW"`
		YourReference       string `json:"YourReference"`
		ZipCode             string `json:"ZipCode"`
	} `json:"Supplier"`
}

type UpdateSupplierResp struct {
	Supplier struct {
		Url                 string `json:"@url"`
		Active              bool   `json:"Active"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Bank                string `json:"Bank"`
		BankAccountNumber   string `json:"BankAccountNumber"`
		BG                  string `json:"BG"`
		BIC                 string `json:"BIC"`
		BranchCode          string `json:"BranchCode"`
		City                string `json:"City"`
		ClearingNumber      string `json:"ClearingNumber"`
		Comments            string `json:"Comments"`
		CostCenter          string `json:"CostCenter"`
		Country             string `json:"Country"`
		CountryCode         string `json:"CountryCode"`
		Currency            string `json:"Currency"`
		DisablePaymentFile  bool   `json:"DisablePaymentFile"`
		Email               string `json:"Email"`
		Fax                 string `json:"Fax"`
		IBAN                string `json:"IBAN"`
		Name                string `json:"Name"`
		OrganisationNumber  string `json:"OrganisationNumber"`
		OurReference        string `json:"OurReference"`
		OurCustomerNumber   string `json:"OurCustomerNumber"`
		PG                  string `json:"PG"`
		Phone1              string `json:"Phone1"`
		Phone2              string `json:"Phone2"`
		PreDefinedAccount   string `json:"PreDefinedAccount"`
		Project             string `json:"Project"`
		SupplierNumber      string `json:"SupplierNumber"`
		TermsOfPayment      string `json:"TermsOfPayment"`
		VATNumber           string `json:"VATNumber"`
		VATType             string `json:"VATType"`
		VisitingAddress     string `json:"VisitingAddress"`
		VisitingCity        string `json:"VisitingCity"`
		VisitingCountry     string `json:"VisitingCountry"`
		VisitingCountryCode string `json:"VisitingCountryCode"`
		VisitingZipCode     string `json:"VisitingZipCode"`
		WorkPlace           string `json:"WorkPlace"`
		WWW                 string `json:"WWW"`
		YourReference       string `json:"YourReference"`
		ZipCode             string `json:"ZipCode"`
	} `json:"Supplier"`
}
