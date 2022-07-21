package client

import (
	"context"
	"fmt"
)

const (
	supplierInvoicePaymentsURI = "supplierinvoicepayments"
)

// GetAllSupplierInvoicePayments does _GET
func (c *Client) GetAllSupplierInvoicePayments(ctx context.Context) (*GetAllSupplierInvoicePaymentsResp, error) {
	resp := &GetAllSupplierInvoicePaymentsResp{}

	err := c._GET(ctx, supplierInvoicePaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSupplierInvoicePayment does _POST
func (c *Client) CreateSupplierInvoicePayment(
	ctx context.Context,
	req *CreateSupplierInvoicePaymentReq) (*CreateSupplierInvoicePaymentResp, error) {

	resp := &CreateSupplierInvoicePaymentResp{}

	err := c._POST(ctx, supplierInvoicePaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSupplierInvoicePayment does _GET
func (c *Client) GetSupplierInvoicePayment(ctx context.Context, number int) (*GetSupplierInvoicePaymentResp, error) {
	resp := &GetSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSupplierInvoicePayment does _PUT
func (c *Client) UpdateSupplierInvoicePayment(
	ctx context.Context,
	number int,
	req *UpdateSupplierInvoicePaymentReq) (*UpdateSupplierInvoicePaymentResp, error) {

	resp := &UpdateSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveSupplierInvoicePayment does _DELETE
func (c *Client) RemoveSupplierInvoicePayment(ctx context.Context, number int) error {
	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)
	return c._DELETE(ctx, uri)
}

// BookKeepSupplierInvoicePayment does _PUT
func (c *Client) BookKeepSupplierInvoicePayment(
	ctx context.Context,
	number int) (*BookKeepSupplierInvoicePaymentResp, error) {

	resp := &BookKeepSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllSupplierInvoicePaymentsResp struct {
	SupplierInvoicePayments []struct {
		Url           string `json:"@url"`
		Amount        int    `json:"Amount"`
		Booked        bool   `json:"Booked"`
		Currency      string `json:"Currency"`
		CurrencyRate  int    `json:"CurrencyRate"`
		CurrencyUnit  int    `json:"CurrencyUnit"`
		InvoiceNumber string `json:"InvoiceNumber"`
		Number        int    `json:"Number"`
		PaymentDate   string `json:"PaymentDate"`
		Source        string `json:"Source"`
	} `json:"SupplierInvoicePayments"`
}

type CreateSupplierInvoicePaymentReq struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}

type CreateSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}

type GetSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}

type UpdateSupplierInvoicePaymentReq struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}

type UpdateSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}

type BookKeepSupplierInvoicePaymentResp struct {
	SupplierInvoicePayment struct {
		Url                   string `json:"@url"`
		Amount                int    `json:"Amount"`
		AmountCurrency        int    `json:"AmountCurrency"`
		Booked                bool   `json:"Booked"`
		Currency              string `json:"Currency"`
		CurrencyRate          int    `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		Information           string `json:"Information"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		InvoiceDueDate        string `json:"InvoiceDueDate"`
		InvoiceOCR            string `json:"InvoiceOCR"`
		InvoiceSupplierName   string `json:"InvoiceSupplierName"`
		InvoiceSupplierNumber string `json:"InvoiceSupplierNumber"`
		InvoiceTotal          string `json:"InvoiceTotal"`
		ModeOfPayment         string `json:"ModeOfPayment"`
		Number                int    `json:"Number"`
		PaymentDate           string `json:"PaymentDate"`
		Source                string `json:"Source"`
		VoucherNumber         int    `json:"VoucherNumber"`
		VoucherSeries         string `json:"VoucherSeries"`
		VoucherYear           int    `json:"VoucherYear"`
		WriteOffs             []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"SupplierInvoicePayment"`
}
