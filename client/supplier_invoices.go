package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	supplierInvoiceURI = "supplierinvoices"
)

const (
	getAllSupplierInvoicesFilterParamName = "filter"
)

// GetAllSupplierInvoices does _GET https://api.fortnox.se/3/supplierinvoices/
//
// filter - Enum: "cancelled" "fullypaid" "unpaid" "unpaidoverdue" "unbooked" "pendingpayment" "authorizepending"
// possibility to filter supplier invoices
func (c *Client) GetAllSupplierInvoices(
	ctx context.Context,
	filter *GetAllSupplierInvoicesFilter) (*GetAllSupplierInvoicesResp, error) {

	resp := &GetAllSupplierInvoicesResp{}

	params := filter.urlValues()

	err := c._GET(ctx, supplierInvoiceURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSupplierInvoice does _POST https://api.fortnox.se/3/supplierinvoices/
//
// req - supplier invoice to create
func (c *Client) CreateSupplierInvoice(
	ctx context.Context,
	req *CreateSupplierInvoiceReq) (*CreateSupplierInvoiceResp, error) {

	resp := &CreateSupplierInvoiceResp{}

	err := c._POST(ctx, supplierInvoiceURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSupplierInvoice does _GET https://api.fortnox.se/3/supplierinvoices/{GivenNumber}
//
// givenNumber - identifies the invoice
func (c *Client) GetSupplierInvoice(ctx context.Context, givenNumber int) (*GetSupplierInvoiceResp, error) {
	resp := &GetSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceURI, givenNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}
//
// givenNumber - identifies the invoice
//
// req - supplier invoice to update
func (c *Client) UpdateSupplierInvoice(
	ctx context.Context,
	givenNumber int,
	req *UpdateSupplierInvoiceReq) (*UpdateSupplierInvoiceResp, error) {

	resp := &UpdateSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BookKeepSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/bookkeep
//
// givenNumber - identifies the invoice
func (c *Client) BookKeepSupplierInvoice(ctx context.Context, givenNumber int) (*BookKeepSupplierInvoiceResp, error) {
	resp := &BookKeepSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d/bookkeep", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CancelSupplierInvoice does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/cancel
//
// givenNumber - identifies the invoice
func (c *Client) CancelSupplierInvoice(ctx context.Context, givenNumber int) (*CancelSupplierInvoiceResp, error) {
	resp := &CancelSupplierInvoiceResp{}

	uri := fmt.Sprintf("%s/%d/cancel", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreditSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/credit
//
// givenNumber - identifies the invoice
func (c *Client) CreditSupplierInvoicePayment(
	ctx context.Context,
	givenNumber int) (*CreditSupplierInvoicePaymentResp, error) {

	resp := &CreditSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d/credit", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ApprovalSupplierInvoicePayment does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/approvalpayment
func (c *Client) ApprovalSupplierInvoicePayment(
	ctx context.Context,
	givenNumber int) (*ApprovalSupplierInvoicePaymentResp, error) {

	resp := &ApprovalSupplierInvoicePaymentResp{}

	uri := fmt.Sprintf("%s/%d/approvalpayment", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ApprovalSupplierInvoiceBookKeep does _PUT https://api.fortnox.se/3/supplierinvoices/{GivenNumber}/approvalbookkeep
//
// givenNumber - identifies the invoice
func (c *Client) ApprovalSupplierInvoiceBookKeep(
	ctx context.Context,
	givenNumber int) (*ApprovalSupplierInvoiceBookKeepResp, error) {

	resp := &ApprovalSupplierInvoiceBookKeepResp{}

	uri := fmt.Sprintf("%s/%d/approvalbookkeep", supplierInvoiceURI, givenNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllSupplierInvoicesFilter string

const (
	CancelledSupplierInvoiceFilter        GetAllSupplierInvoicesFilter = "cancelled"
	fullyPaidSupplierInvoiceFilter        GetAllSupplierInvoicesFilter = "fullypaid"
	UnpaidSupplierInvoiceFilter           GetAllSupplierInvoicesFilter = "unpaid"
	UnPaidOverdueSupplierInvoiceFilter    GetAllSupplierInvoicesFilter = "unpaidoverdue"
	UnbookedSupplierInvoiceFilter         GetAllSupplierInvoicesFilter = "unbooked"
	PendingPaymentSupplierInvoiceFilter   GetAllSupplierInvoicesFilter = "pendingpayment"
	AuthorizePendingSupplierInvoiceFilter GetAllSupplierInvoicesFilter = "authorizepending"
)

func (f *GetAllSupplierInvoicesFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	fStr := string(*f)
	if strings.TrimSpace(fStr) != "" {
		params[getAllSupplierInvoicesFilterParamName] = []string{fStr}
	}

	return params
}

type GetAllSupplierInvoicesResp struct {
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

type CreateSupplierInvoiceReq struct {
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

type CreateSupplierInvoiceResp struct {
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

type GetSupplierInvoiceResp struct {
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

type UpdateSupplierInvoiceReq struct {
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

type UpdateSupplierInvoiceResp struct {
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

type BookKeepSupplierInvoiceResp struct {
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

type CancelSupplierInvoiceResp struct {
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

type CreditSupplierInvoicePaymentResp struct {
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

type ApprovalSupplierInvoicePaymentResp struct {
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

type ApprovalSupplierInvoiceBookKeepResp struct {
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
