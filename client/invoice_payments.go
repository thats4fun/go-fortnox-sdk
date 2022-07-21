package client

import (
	"context"
	"fmt"
)

const (
	invoicePaymentsURI = "invoicepayments"
)

// GetAllInvoicePayments does _GET https://api.fortnox.se/3/invoicepayments/
func (c *Client) GetAllInvoicePayments(ctx context.Context) (*GetAllInvoicePaymentsResp, error) {
	resp := &GetAllInvoicePaymentsResp{}

	err := c._GET(ctx, invoicePaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateInvoicePayment does _POST https://api.fortnox.se/3/invoicepayments/
//
// req - invoice payment to create
func (c *Client) CreateInvoicePayment(
	ctx context.Context,
	req *CreateInvoicePaymentReq) (*CreateInvoicePaymentResp, error) {

	resp := &CreateInvoicePaymentResp{}

	err := c._POST(ctx, invoicePaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInvoicePayment does _GET https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
func (c *Client) GetInvoicePayment(ctx context.Context, number string) (*GetInvoicePaymentResp, error) {
	resp := &GetInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateInvoicePayment does _PUT https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
//
// req - invoice payment to update
func (c *Client) UpdateInvoicePayment(
	ctx context.Context,
	number string, req *UpdateInvoicePaymentReq) (*UpdateInvoicePaymentResp, error) {

	resp := &UpdateInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RemoveInvoicePayment does _DELETE https://api.fortnox.se/3/invoicepayments/{Number}
//
// number - identifies the invoice payment
func (c *Client) RemoveInvoicePayment(ctx context.Context, number string) error {
	uri := fmt.Sprintf("%s/%s", invoicePaymentsURI, number)
	return c._DELETE(ctx, uri)
}

// BookKeepInvoicePayment does _PUT https://api.fortnox.se/3/invoicepayments/{Number}/bookkeep
//
// number - identifies the invoice payment
func (c *Client) BookKeepInvoicePayment(
	ctx context.Context,
	number string,
	req *BookKeepInvoicePaymentReq) (*BookKeepInvoicePaymentResp, error) {

	resp := &BookKeepInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%s/bookkeep", invoicePaymentsURI, number)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllInvoicePaymentsResp struct {
	InvoicePayments []struct {
		Url           string `json:"@url"`
		Amount        int    `json:"Amount"`
		Booked        bool   `json:"Booked"`
		Currency      string `json:"Currency"`
		CurrencyRate  int    `json:"CurrencyRate"`
		CurrencyUnit  int    `json:"CurrencyUnit"`
		InvoiceNumber int    `json:"InvoiceNumber"`
		Number        string `json:"Number"`
		PaymentDate   string `json:"PaymentDate"`
		Source        string `json:"Source"`
	} `json:"InvoicePayments"`
}

type CreateInvoicePaymentReq struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type CreateInvoicePaymentResp struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type BookKeepInvoicePaymentReq struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type BookKeepInvoicePaymentResp struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type GetInvoicePaymentResp struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type UpdateInvoicePaymentReq struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}

type UpdateInvoicePaymentResp struct {
	InvoicePayment struct {
		Url                       string `json:"@url"`
		Amount                    int    `json:"Amount"`
		AmountCurrency            int    `json:"AmountCurrency"`
		Booked                    bool   `json:"Booked"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceCustomerName       string `json:"InvoiceCustomerName"`
		InvoiceCustomerNumber     string `json:"InvoiceCustomerNumber"`
		InvoiceNumber             int    `json:"InvoiceNumber"`
		InvoiceDueDate            string `json:"InvoiceDueDate"`
		InvoiceOCR                string `json:"InvoiceOCR"`
		InvoiceTotal              string `json:"InvoiceTotal"`
		ModeOfPayment             string `json:"ModeOfPayment"`
		ModeOfPaymentAccount      int    `json:"ModeOfPaymentAccount"`
		Number                    string `json:"Number"`
		PaymentDate               string `json:"PaymentDate"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		Source                    string `json:"Source"`
		WriteOffs                 []struct {
			Amount                 int    `json:"Amount"`
			AccountNumber          int    `json:"AccountNumber"`
			CostCenter             string `json:"CostCenter"`
			Currency               string `json:"Currency"`
			Description            string `json:"Description"`
			TransactionInformation string `json:"TransactionInformation"`
			Project                string `json:"Project"`
		} `json:"WriteOffs"`
	} `json:"InvoicePayment"`
}
