package client

import (
	"context"
	"fmt"
)

const (
	customersURI = "customers"
)

// GetAllCustomers does _GET https://api.fortnox.se/3/customers/
func (c *Client) GetAllCustomers(ctx context.Context) ([]Customer, error) {
	resp := &GetAllCustomersResp{}

	err := c._GET(ctx, customersURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Customers, nil
}

// CreateCustomer does _POST https://api.fortnox.se/3/customers/
//
// req - customer to create
func (c *Client) CreateCustomer(ctx context.Context, cus *Customer) (*Customer, error) {
	req := &CreateCustomerReq{Customer: *cus}
	resp := &CreateCustomerResp{}

	err := c._POST(ctx, customersURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// GetCustomer does _GET https://api.fortnox.se/3/customers/{CustomerId}
//
// customerNumber - identifies the customer
func (c *Client) GetCustomer(ctx context.Context, customerNumber string) (*Customer, error) {
	resp := &GetCustomerResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// UpdateCustomer does _PUT https://api.fortnox.se/3/employees/{CustomerId}
//
// customerNumber - identifies the customer
//
// req - customer to update
func (c *Client) UpdateCustomer(
	ctx context.Context,
	customerNumber string,
	cus *Customer) (*Customer, error) {

	req := &UpdateCustomerReq{Customer: *cus}
	resp := &UpdateCustomerResp{}

	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Customer, nil
}

// DeleteCustomer does _DELETE https://api.fortnox.se/3/customers/{CustomerNumber}
//
// customerNumber - identifies the customer
func (c *Client) DeleteCustomer(ctx context.Context, customerNumber string) error {
	uri := fmt.Sprintf("%s/%s", customersURI, customerNumber)
	return c._DELETE(ctx, uri)
}

type GetAllCustomersResp struct {
	Customers []Customer `json:"Customers"`
}

type CreateCustomerReq struct {
	Customer Customer `json:"Customer"`
}

type CreateCustomerResp struct {
	Customer Customer `json:"Customer"`
}

type GetCustomerResp struct {
	Customer Customer `json:"Customer"`
}

type UpdateCustomerReq struct {
	Customer Customer `json:"Customer"`
}

type UpdateCustomerResp struct {
	Customer Customer `json:"Customer"`
}

type Customer struct {
	Url                      string               `json:"@url,omitempty"`
	Address1                 string               `json:"Address1,omitempty"`
	Address2                 string               `json:"Address2,omitempty"`
	City                     string               `json:"City,omitempty"`
	Country                  string               `json:"Country,omitempty"`
	Comments                 string               `json:"Comments,omitempty"`
	Currency                 string               `json:"Currency,omitempty"`
	CostCenter               string               `json:"CostCenter,omitempty"`
	CountryCode              string               `json:"CountryCode,omitempty"`
	Active                   bool                 `json:"Active,omitempty"`
	CustomerNumber           string               `json:"CustomerNumber,omitempty"`
	DefaultDeliveryTypes     DefaultDeliveryTypes `json:"DefaultDeliveryTypes,omitempty"`
	DefaultTemplates         DefaultTemplates     `json:"DefaultTemplates,omitempty"`
	DeliveryAddress1         string               `json:"DeliveryAddress1,omitempty"`
	DeliveryAddress2         string               `json:"DeliveryAddress2,omitempty"`
	DeliveryCity             string               `json:"DeliveryCity,omitempty"`
	DeliveryCountry          string               `json:"DeliveryCountry,omitempty"`
	DeliveryCountryCode      string               `json:"DeliveryCountryCode,omitempty"`
	DeliveryFax              string               `json:"DeliveryFax,omitempty"`
	DeliveryName             string               `json:"DeliveryName,omitempty"`
	DeliveryPhone1           string               `json:"DeliveryPhone1,omitempty"`
	DeliveryPhone2           string               `json:"DeliveryPhone2,omitempty"`
	DeliveryZipCode          string               `json:"DeliveryZipCode,omitempty"`
	Email                    string               `json:"Email,omitempty"`
	EmailInvoice             string               `json:"EmailInvoice,omitempty"`
	EmailInvoiceBCC          string               `json:"EmailInvoiceBCC,omitempty"`
	EmailInvoiceCC           string               `json:"EmailInvoiceCC,omitempty"`
	EmailOffer               string               `json:"EmailOffer,omitempty"`
	EmailOfferBCC            string               `json:"EmailOfferBCC,omitempty"`
	EmailOfferCC             string               `json:"EmailOfferCC,omitempty"`
	EmailOrder               string               `json:"EmailOrder,omitempty"`
	EmailOrderBCC            string               `json:"EmailOrderBCC,omitempty"`
	EmailOrderCC             string               `json:"EmailOrderCC,omitempty"`
	ExternalReference        string               `json:"ExternalReference,omitempty"`
	Fax                      string               `json:"Fax,omitempty"`
	GLN                      string               `json:"GLN,omitempty"`
	GLNDelivery              string               `json:"GLNDelivery,omitempty"`
	InvoiceAdministrationFee string               `json:"InvoiceAdministrationFee,omitempty"`
	InvoiceDiscount          int                  `json:"InvoiceDiscount,omitempty"`
	InvoiceFreight           string               `json:"InvoiceFreight,omitempty"`
	InvoiceRemark            string               `json:"InvoiceRemark,omitempty"`
	Name                     string               `json:"Name,omitempty"`
	OrganisationNumber       string               `json:"OrganisationNumber,omitempty"`
	OurReference             string               `json:"OurReference,omitempty"`
	Phone1                   string               `json:"Phone1,omitempty"`
	Phone2                   string               `json:"Phone2,omitempty"`
	PriceList                string               `json:"PriceList,omitempty"`
	Project                  string               `json:"Project,omitempty"`
	SalesAccount             string               `json:"SalesAccount,omitempty"`
	ShowPriceVATIncluded     bool                 `json:"ShowPriceVATIncluded,omitempty"`
	TermsOfDelivery          string               `json:"TermsOfDelivery,omitempty"`
	TermsOfPayment           string               `json:"TermsOfPayment,omitempty"`
	Type                     string               `json:"Type,omitempty"`
	VATNumber                string               `json:"VATNumber,omitempty"`
	VATType                  string               `json:"VATType,omitempty"`
	VisitingAddress          string               `json:"VisitingAddress,omitempty"`
	VisitingCity             string               `json:"VisitingCity,omitempty"`
	VisitingCountry          string               `json:"VisitingCountry,omitempty"`
	VisitingCountryCode      string               `json:"VisitingCountryCode,omitempty"`
	VisitingZipCode          string               `json:"VisitingZipCode,omitempty"`
	WayOfDelivery            string               `json:"WayOfDelivery,omitempty"`
	WWW                      string               `json:"WWW,omitempty"`
	YourReference            string               `json:"YourReference,omitempty"`
	ZipCode                  string               `json:"ZipCode,omitempty"`
}
type DefaultDeliveryTypes struct {
	Invoice string `json:"Invoice,omitempty"`
	Order   string `json:"Order,omitempty"`
	Offer   string `json:"Offer,omitempty"`
}

type DefaultTemplates struct {
	CashInvoice string `json:"CashInvoice,omitempty"`
	Invoice     string `json:"Invoice,omitempty"`
	Offer       string `json:"Offer,omitempty"`
	Order       string `json:"Order,omitempty"`
}
