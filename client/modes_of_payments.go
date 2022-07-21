package client

import (
	"context"
	"fmt"
)

const (
	modesOfPaymentsURI = "modesofpayments"
)

// GetAllModesOfPayments does _GET https://api.fortnox.se/3/supplierinvoices/
func (c *Client) GetAllModesOfPayments(ctx context.Context) (*GetAllModesOfPaymentsResp, error) {
	resp := &GetAllModesOfPaymentsResp{}

	err := c._GET(ctx, modesOfPaymentsURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateModeOfPayment does _POST https://api.fortnox.se/3/supplierinvoices/
//
// req - mode of payment to create
func (c *Client) CreateModeOfPayment(
	ctx context.Context,
	req *CreateModeOfPaymentReq) (*CreateModeOfPaymentResp, error) {

	resp := &CreateModeOfPaymentResp{}

	err := c._POST(ctx, modesOfPaymentsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetModeOfPayment does _GET https://api.fortnox.se/3/supplierinvoices/{GivenNumber}
//
// code - identifies the mode of payment
func (c *Client) GetModeOfPayment(ctx context.Context, code string) (*GetModeOfPaymentResp, error) {
	resp := &GetModeOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", modesOfPaymentsURI, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateModeOfPayment does _PUT https://api.fortnox.se/3/modesofpayments/{Code}
//
// code - identifies the mode of payment
//
// req - mode of payment to update
func (c *Client) UpdateModeOfPayment(
	ctx context.Context,
	code string,
	req *UpdateModeOfPaymentReq) (*UpdateModeOfPaymentResp, error) {

	resp := &UpdateModeOfPaymentResp{}

	uri := fmt.Sprintf("%s/%s", modesOfPaymentsURI, code)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllModesOfPaymentsResp struct {
	SupplierInvoices []struct {
		Url                   string `json:"@url"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancel                bool   `json:"Cancel"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		Project               string `json:"Project"`
		SupplierNumber        string `json:"SupplierNumber"`
		SupplierName          string `json:"SupplierName"`
		Total                 string `json:"Total"`
		AuthorizerName        string `json:"AuthorizerName"`
		Vouchers              []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoices"`
}

type CreateModeOfPaymentReq struct {
	SupplierInvoice struct {
		Url                   string `json:"@url"`
		AdministrationFee     string `json:"AdministrationFee"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancelled             bool   `json:"Cancelled"`
		Comments              string `json:"Comments"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		CreditReference       int    `json:"CreditReference"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DisablePaymentFile    bool   `json:"DisablePaymentFile"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		Freight               string `json:"Freight"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		OCR                   string `json:"OCR"`
		OurReference          string `json:"OurReference"`
		PaymentPending        bool   `json:"PaymentPending"`
		Project               string `json:"Project"`
		RoundOffValue         string `json:"RoundOffValue"`
		SupplierInvoiceRows   []struct {
			Account                int    `json:"Account"`
			ArticleNumber          string `json:"ArticleNumber"`
			Code                   string `json:"Code"`
			CostCenter             string `json:"CostCenter"`
			AccountDescription     string `json:"AccountDescription"`
			ItemDescription        string `json:"ItemDescription"`
			Debit                  int    `json:"Debit"`
			DebitCurrency          int    `json:"DebitCurrency"`
			Credit                 int    `json:"Credit"`
			CreditCurrency         int    `json:"CreditCurrency"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
			Price                  int    `json:"Price"`
			Quantity               int    `json:"Quantity"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			StockPointCode         string `json:"StockPointCode"`
			StockLocationCode      string `json:"StockLocationCode"`
		} `json:"SupplierInvoiceRows"`
		SupplierNumber   string `json:"SupplierNumber"`
		SupplierName     string `json:"SupplierName"`
		Total            string `json:"Total"`
		VAT              string `json:"VAT"`
		YourReference    string `json:"YourReference"`
		VoucherNumber    int    `json:"VoucherNumber"`
		VoucherSeries    string `json:"VoucherSeries"`
		VoucherYear      int    `json:"VoucherYear"`
		VATType          string `json:"VATType"`
		SalesType        string `json:"SalesType"`
		AccountingMethod string `json:"AccountingMethod"`
		Vouchers         []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoice"`
}

type CreateModeOfPaymentResp struct {
	SupplierInvoice struct {
		Url                   string `json:"@url"`
		AdministrationFee     string `json:"AdministrationFee"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancelled             bool   `json:"Cancelled"`
		Comments              string `json:"Comments"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		CreditReference       int    `json:"CreditReference"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DisablePaymentFile    bool   `json:"DisablePaymentFile"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		Freight               string `json:"Freight"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		OCR                   string `json:"OCR"`
		OurReference          string `json:"OurReference"`
		PaymentPending        bool   `json:"PaymentPending"`
		Project               string `json:"Project"`
		RoundOffValue         string `json:"RoundOffValue"`
		SupplierInvoiceRows   []struct {
			Account                int    `json:"Account"`
			ArticleNumber          string `json:"ArticleNumber"`
			Code                   string `json:"Code"`
			CostCenter             string `json:"CostCenter"`
			AccountDescription     string `json:"AccountDescription"`
			ItemDescription        string `json:"ItemDescription"`
			Debit                  int    `json:"Debit"`
			DebitCurrency          int    `json:"DebitCurrency"`
			Credit                 int    `json:"Credit"`
			CreditCurrency         int    `json:"CreditCurrency"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
			Price                  int    `json:"Price"`
			Quantity               int    `json:"Quantity"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			StockPointCode         string `json:"StockPointCode"`
			StockLocationCode      string `json:"StockLocationCode"`
		} `json:"SupplierInvoiceRows"`
		SupplierNumber   string `json:"SupplierNumber"`
		SupplierName     string `json:"SupplierName"`
		Total            string `json:"Total"`
		VAT              string `json:"VAT"`
		YourReference    string `json:"YourReference"`
		VoucherNumber    int    `json:"VoucherNumber"`
		VoucherSeries    string `json:"VoucherSeries"`
		VoucherYear      int    `json:"VoucherYear"`
		VATType          string `json:"VATType"`
		SalesType        string `json:"SalesType"`
		AccountingMethod string `json:"AccountingMethod"`
		Vouchers         []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoice"`
}

type GetModeOfPaymentResp struct {
	SupplierInvoice struct {
		Url                   string `json:"@url"`
		AdministrationFee     string `json:"AdministrationFee"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancelled             bool   `json:"Cancelled"`
		Comments              string `json:"Comments"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		CreditReference       int    `json:"CreditReference"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DisablePaymentFile    bool   `json:"DisablePaymentFile"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		Freight               string `json:"Freight"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		OCR                   string `json:"OCR"`
		OurReference          string `json:"OurReference"`
		PaymentPending        bool   `json:"PaymentPending"`
		Project               string `json:"Project"`
		RoundOffValue         string `json:"RoundOffValue"`
		SupplierInvoiceRows   []struct {
			Account                int    `json:"Account"`
			ArticleNumber          string `json:"ArticleNumber"`
			Code                   string `json:"Code"`
			CostCenter             string `json:"CostCenter"`
			AccountDescription     string `json:"AccountDescription"`
			ItemDescription        string `json:"ItemDescription"`
			Debit                  int    `json:"Debit"`
			DebitCurrency          int    `json:"DebitCurrency"`
			Credit                 int    `json:"Credit"`
			CreditCurrency         int    `json:"CreditCurrency"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
			Price                  int    `json:"Price"`
			Quantity               int    `json:"Quantity"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			StockPointCode         string `json:"StockPointCode"`
			StockLocationCode      string `json:"StockLocationCode"`
		} `json:"SupplierInvoiceRows"`
		SupplierNumber   string `json:"SupplierNumber"`
		SupplierName     string `json:"SupplierName"`
		Total            string `json:"Total"`
		VAT              string `json:"VAT"`
		YourReference    string `json:"YourReference"`
		VoucherNumber    int    `json:"VoucherNumber"`
		VoucherSeries    string `json:"VoucherSeries"`
		VoucherYear      int    `json:"VoucherYear"`
		VATType          string `json:"VATType"`
		SalesType        string `json:"SalesType"`
		AccountingMethod string `json:"AccountingMethod"`
		Vouchers         []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoice"`
}

type UpdateModeOfPaymentReq struct {
	SupplierInvoice struct {
		Url                   string `json:"@url"`
		AdministrationFee     string `json:"AdministrationFee"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancelled             bool   `json:"Cancelled"`
		Comments              string `json:"Comments"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		CreditReference       int    `json:"CreditReference"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DisablePaymentFile    bool   `json:"DisablePaymentFile"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		Freight               string `json:"Freight"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		OCR                   string `json:"OCR"`
		OurReference          string `json:"OurReference"`
		PaymentPending        bool   `json:"PaymentPending"`
		Project               string `json:"Project"`
		RoundOffValue         string `json:"RoundOffValue"`
		SupplierInvoiceRows   []struct {
			Account                int    `json:"Account"`
			ArticleNumber          string `json:"ArticleNumber"`
			Code                   string `json:"Code"`
			CostCenter             string `json:"CostCenter"`
			AccountDescription     string `json:"AccountDescription"`
			ItemDescription        string `json:"ItemDescription"`
			Debit                  int    `json:"Debit"`
			DebitCurrency          int    `json:"DebitCurrency"`
			Credit                 int    `json:"Credit"`
			CreditCurrency         int    `json:"CreditCurrency"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
			Price                  int    `json:"Price"`
			Quantity               int    `json:"Quantity"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			StockPointCode         string `json:"StockPointCode"`
			StockLocationCode      string `json:"StockLocationCode"`
		} `json:"SupplierInvoiceRows"`
		SupplierNumber   string `json:"SupplierNumber"`
		SupplierName     string `json:"SupplierName"`
		Total            string `json:"Total"`
		VAT              string `json:"VAT"`
		YourReference    string `json:"YourReference"`
		VoucherNumber    int    `json:"VoucherNumber"`
		VoucherSeries    string `json:"VoucherSeries"`
		VoucherYear      int    `json:"VoucherYear"`
		VATType          string `json:"VATType"`
		SalesType        string `json:"SalesType"`
		AccountingMethod string `json:"AccountingMethod"`
		Vouchers         []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoice"`
}

type UpdateModeOfPaymentResp struct {
	SupplierInvoice struct {
		Url                   string `json:"@url"`
		AdministrationFee     string `json:"AdministrationFee"`
		Balance               string `json:"Balance"`
		Booked                bool   `json:"Booked"`
		Cancelled             bool   `json:"Cancelled"`
		Comments              string `json:"Comments"`
		CostCenter            string `json:"CostCenter"`
		Credit                bool   `json:"Credit"`
		CreditReference       int    `json:"CreditReference"`
		Currency              string `json:"Currency"`
		CurrencyRate          string `json:"CurrencyRate"`
		CurrencyUnit          int    `json:"CurrencyUnit"`
		DisablePaymentFile    bool   `json:"DisablePaymentFile"`
		DueDate               string `json:"DueDate"`
		ExternalInvoiceNumber string `json:"ExternalInvoiceNumber"`
		ExternalInvoiceSeries string `json:"ExternalInvoiceSeries"`
		Freight               string `json:"Freight"`
		GivenNumber           string `json:"GivenNumber"`
		InvoiceDate           string `json:"InvoiceDate"`
		InvoiceNumber         string `json:"InvoiceNumber"`
		OCR                   string `json:"OCR"`
		OurReference          string `json:"OurReference"`
		PaymentPending        bool   `json:"PaymentPending"`
		Project               string `json:"Project"`
		RoundOffValue         string `json:"RoundOffValue"`
		SupplierInvoiceRows   []struct {
			Account                int    `json:"Account"`
			ArticleNumber          string `json:"ArticleNumber"`
			Code                   string `json:"Code"`
			CostCenter             string `json:"CostCenter"`
			AccountDescription     string `json:"AccountDescription"`
			ItemDescription        string `json:"ItemDescription"`
			Debit                  int    `json:"Debit"`
			DebitCurrency          int    `json:"DebitCurrency"`
			Credit                 int    `json:"Credit"`
			CreditCurrency         int    `json:"CreditCurrency"`
			Project                string `json:"Project"`
			TransactionInformation string `json:"TransactionInformation"`
			Price                  int    `json:"Price"`
			Quantity               int    `json:"Quantity"`
			Total                  int    `json:"Total"`
			Unit                   string `json:"Unit"`
			StockPointCode         string `json:"StockPointCode"`
			StockLocationCode      string `json:"StockLocationCode"`
		} `json:"SupplierInvoiceRows"`
		SupplierNumber   string `json:"SupplierNumber"`
		SupplierName     string `json:"SupplierName"`
		Total            string `json:"Total"`
		VAT              string `json:"VAT"`
		YourReference    string `json:"YourReference"`
		VoucherNumber    int    `json:"VoucherNumber"`
		VoucherSeries    string `json:"VoucherSeries"`
		VoucherYear      int    `json:"VoucherYear"`
		VATType          string `json:"VATType"`
		SalesType        string `json:"SalesType"`
		AccountingMethod string `json:"AccountingMethod"`
		Vouchers         []struct {
			Number        int    `json:"Number"`
			Year          int    `json:"Year"`
			Series        string `json:"Series"`
			ReferenceType string `json:"ReferenceType"`
		} `json:"Vouchers"`
		FinalPayDate string `json:"FinalPayDate"`
	} `json:"SupplierInvoice"`
}
