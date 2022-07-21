package client

import (
	"context"
	"fmt"
)

const (
	customersURI = "customers"
)

// GetAllCustomers does _GET https://api.fortnox.se/3/customers/
func (c *Client) GetAllCustomers(ctx context.Context) (*GetAllCustomersResp, error) {
	resp := &GetAllCustomersResp{}

	err := c._GET(ctx, customersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateCustomer does _POST https://api.fortnox.se/3/customers/
//
// req - customer to create
func (c *Client) CreateCustomer(ctx context.Context, req *CreateCustomerReq) (*CreateCustomerResp, error) {
	resp := &CreateCustomerResp{}

	err := c._POST(ctx, customersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetCustomer does _GET https://api.fortnox.se/3/customers/{CustomerId}
//
// customerNumber - identifies the customer
func (c *Client) GetCustomer(ctx context.Context, customerNumber string) (*GetCustomerResp, error) {
	resp := &GetCustomerResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateCustomer does _PUT https://api.fortnox.se/3/employees/{CustomerId}
//
// customerNumber - identifies the customer
//
// req - customer to update
func (c *Client) UpdateCustomer(
	ctx context.Context,
	customerNumber string,
	req *UpdateCustomerReq) (*UpdateCustomerResp, error) {

	resp := &UpdateCustomerResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteCustomer does _DELETE https://api.fortnox.se/3/customers/{CustomerNumber}
//
// customerNumber - identifies the customer
func (c *Client) DeleteCustomer(ctx context.Context, customerNumber string) error {
	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)
	return c._DELETE(ctx, uri)
}

type GetAllCustomersResp struct {
	Customers []struct {
		Url                string `json:"@url"`
		Address1           string `json:"Address1"`
		Address2           string `json:"Address2"`
		City               string `json:"City"`
		CustomerNumber     string `json:"CustomerNumber"`
		Email              string `json:"Email"`
		Name               string `json:"Name"`
		OrganisationNumber string `json:"OrganisationNumber"`
		Phone              string `json:"Phone"`
		ZipCode            string `json:"ZipCode"`
	} `json:"Customers"`
}

type CreateCustomerReq struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}

type CreateCustomerResp struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}

type GetCustomerResp struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}

type UpdateCustomerReq struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}

type UpdateCustomerResp struct {
	Customer struct {
		Url                  string `json:"@url"`
		Address1             string `json:"Address1"`
		Address2             string `json:"Address2"`
		City                 string `json:"City"`
		Country              string `json:"Country"`
		Comments             string `json:"Comments"`
		Currency             string `json:"Currency"`
		CostCenter           string `json:"CostCenter"`
		CountryCode          string `json:"CountryCode"`
		Active               bool   `json:"Active"`
		CustomerNumber       string `json:"CustomerNumber"`
		DefaultDeliveryTypes struct {
			Invoice string `json:"Invoice"`
			Order   string `json:"Order"`
			Offer   string `json:"Offer"`
		} `json:"DefaultDeliveryTypes"`
		DefaultTemplates struct {
			CashInvoice string `json:"CashInvoice"`
			Invoice     string `json:"Invoice"`
			Offer       string `json:"Offer"`
			Order       string `json:"Order"`
		} `json:"DefaultTemplates"`
		DeliveryAddress1         string `json:"DeliveryAddress1"`
		DeliveryAddress2         string `json:"DeliveryAddress2"`
		DeliveryCity             string `json:"DeliveryCity"`
		DeliveryCountry          string `json:"DeliveryCountry"`
		DeliveryCountryCode      string `json:"DeliveryCountryCode"`
		DeliveryFax              string `json:"DeliveryFax"`
		DeliveryName             string `json:"DeliveryName"`
		DeliveryPhone1           string `json:"DeliveryPhone1"`
		DeliveryPhone2           string `json:"DeliveryPhone2"`
		DeliveryZipCode          string `json:"DeliveryZipCode"`
		Email                    string `json:"Email"`
		EmailInvoice             string `json:"EmailInvoice"`
		EmailInvoiceBCC          string `json:"EmailInvoiceBCC"`
		EmailInvoiceCC           string `json:"EmailInvoiceCC"`
		EmailOffer               string `json:"EmailOffer"`
		EmailOfferBCC            string `json:"EmailOfferBCC"`
		EmailOfferCC             string `json:"EmailOfferCC"`
		EmailOrder               string `json:"EmailOrder"`
		EmailOrderBCC            string `json:"EmailOrderBCC"`
		EmailOrderCC             string `json:"EmailOrderCC"`
		ExternalReference        string `json:"ExternalReference"`
		Fax                      string `json:"Fax"`
		GLN                      string `json:"GLN"`
		GLNDelivery              string `json:"GLNDelivery"`
		InvoiceAdministrationFee string `json:"InvoiceAdministrationFee"`
		InvoiceDiscount          int    `json:"InvoiceDiscount"`
		InvoiceFreight           string `json:"InvoiceFreight"`
		InvoiceRemark            string `json:"InvoiceRemark"`
		Name                     string `json:"Name"`
		OrganisationNumber       string `json:"OrganisationNumber"`
		OurReference             string `json:"OurReference"`
		Phone1                   string `json:"Phone1"`
		Phone2                   string `json:"Phone2"`
		PriceList                string `json:"PriceList"`
		Project                  string `json:"Project"`
		SalesAccount             string `json:"SalesAccount"`
		ShowPriceVATIncluded     bool   `json:"ShowPriceVATIncluded"`
		TermsOfDelivery          string `json:"TermsOfDelivery"`
		TermsOfPayment           string `json:"TermsOfPayment"`
		Type                     string `json:"Type"`
		VATNumber                string `json:"VATNumber"`
		VATType                  string `json:"VATType"`
		VisitingAddress          string `json:"VisitingAddress"`
		VisitingCity             string `json:"VisitingCity"`
		VisitingCountry          string `json:"VisitingCountry"`
		VisitingCountryCode      string `json:"VisitingCountryCode"`
		VisitingZipCode          string `json:"VisitingZipCode"`
		WayOfDelivery            string `json:"WayOfDelivery"`
		WWW                      string `json:"WWW"`
		YourReference            string `json:"YourReference"`
		ZipCode                  string `json:"ZipCode"`
	} `json:"Customer"`
}
