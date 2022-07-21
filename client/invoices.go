package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	invoicesURI = "invoices"
)

// GetInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}
//
// documentNumber - identifies the invoice
func (c *Client) GetInvoice(ctx context.Context, documentNumber string) (*GetInvoiceResp, error) {
	resp := &GetInvoiceResp{}

	uri := fmt.Sprintf("%s/%s", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateInvoice odes _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}
//
// documentNumber - identifies the invoice
//
// req - payload
func (c *Client) UpdateInvoice(
	ctx context.Context,
	documentNumber string,
	req *UpdateInvoiceReq) (*UpdateInvoiceResp, error) {

	resp := &UpdateInvoiceResp{}

	uri := fmt.Sprintf("%s/%s", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAllInvoices does _GET https://api.fortnox.se/3/invoices
//
// queryParams - filters
func (c *Client) GetAllInvoices(ctx context.Context, queryParams *GetAllInvoicesQueryParams) (*GetAllInvoicesResp, error) {
	resp := &GetAllInvoicesResp{}

	var params url.Values
	if queryParams != nil {
		params = queryParams.urlValues()
	}

	err := c._GET(ctx, invoicesURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//CreateInvoice does _POST https://api.fortnox.se/3/invoices
//
// req - payload
func (c *Client) CreateInvoice(ctx context.Context, req *CreateInvoiceReq) (*CreateInvoiceResp, error) {
	resp := &CreateInvoiceResp{}

	err := c._POST(ctx, invoicesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BookKeepInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/bookkeep
//
// documentNumber - identifies the invoice
func (c *Client) BookKeepInvoice(ctx context.Context, documentNumber string) (*BookKeepInvoiceResp, error) {
	resp := &BookKeepInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/bookkeep", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CancelInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/cancel
//
// documentNumber - identifies the invoice
func (c *Client) CancelInvoice(ctx context.Context, documentNumber string) (*CancelInvoiceResp, error) {
	resp := &CancelInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/cancel", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreditInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/credit
//
// documentNumber - identifies the invoice
func (c *Client) CreditInvoice(ctx context.Context, documentNumber string) (*CreditInvoiceResp, error) {
	resp := &CreditInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/credit", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetInvoiceAsSent does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/externalprint
//
// documentNumber - identifies the invoice
func (c *Client) SetInvoiceAsSent(ctx context.Context, documentNumber string) (*SetInvoiceAsSentResp, error) {
	resp := &SetInvoiceAsSentResp{}

	uri := fmt.Sprintf("%s/%s/externalprint", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetInvoiceAsDone does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/warehouseready
//
// documentNumber - identifies the invoice
func (c *Client) SetInvoiceAsDone(ctx context.Context, documentNumber string) (*SetInvoiceAsDoneResp, error) {
	resp := &SetInvoiceAsDoneResp{}

	uri := fmt.Sprintf("%s/%s/warehouseready", invoicesURI, documentNumber)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PrintInvoice does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/print
//
// documentNumber - identifies the invoice
func (c *Client) PrintInvoice(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/print", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsEmail does _PUT https://api.fortnox.se/3/invoices/{DocumentNumber}/email
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEmail(ctx context.Context, documentNumber string) (*SendInvoiceAsEmailResp, error) {
	resp := &SendInvoiceAsEmailResp{}

	uri := fmt.Sprintf("%s/%s/email", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsReminder does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/printreminder
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsReminder(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/printreminder", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PreviewInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/preview
//
// documentNumber - identifies the invoice
func (c *Client) PreviewInvoice(ctx context.Context, documentNumber string) (*[]byte, error) {
	resp := &[]byte{}

	uri := fmt.Sprintf("%s/%s/preview", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsEPrint does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/eprint
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEPrint(ctx context.Context, documentNumber string) (*SendInvoiceAsEPrintResp, error) {
	resp := &SendInvoiceAsEPrintResp{}

	uri := fmt.Sprintf("%s/%s/eprint", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendInvoiceAsEInvoice does _GET https://api.fortnox.se/3/invoices/{DocumentNumber}/einvoice
//
// documentNumber - identifies the invoice
func (c *Client) SendInvoiceAsEInvoice(ctx context.Context, documentNumber string) (*SendInvoiceAsEInvoiceResp, error) {
	resp := &SendInvoiceAsEInvoiceResp{}

	uri := fmt.Sprintf("%s/%s/einvoice", invoicesURI, documentNumber)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllInvoicesQueryParams struct {
	Filter                    GetAllInvoicesFilter
	CostCenter                string
	CustomerName              string
	CustomerNumber            string
	Label                     string
	DocumentNumber            string
	FromDate                  string
	ToDate                    string
	FromFinalPayDate          string
	ToFinalPayDate            string
	LastModified              string
	NotCompleted              string
	Ocr                       string
	OurReference              string
	Project                   string
	Sent                      string
	ExternalInvoiceReference1 string
	ExternalInvoiceReference2 string
	YourReference             string
	InvoiceType               string
	ArticleNumber             string
	ArticleDescription        string
	Currency                  string
	AccountNumberFrom         string
	AccountNumberTo           string
	YourOrderNumber           string
	Credit                    string
	SortBy                    GetAllInvoicesSortBy
}

func (p GetAllInvoicesQueryParams) urlValues() url.Values {
	params := url.Values{}

	filterStr := string(p.Filter)
	if strings.TrimSpace(filterStr) != "" {
		params["filter"] = []string{filterStr}
	}

	return params
}

type GetAllInvoicesSortBy string

const (
	CustomerName   = "customername"
	CustomerNumber = "customernumber"
	DocumentNumber = "documentnumber"
	InvoiceDate    = "invoicedate"
	Ocr            = "ocr"
	Total          = "total"
)

type GetAllInvoicesFilter string

const (
	Cancelled     GetAllInvoicesFilter = "cancelled"
	FullyPaid     GetAllInvoicesFilter = "fullypaid"
	Unpaid        GetAllInvoicesFilter = "unpaid"
	UnpaidOverDue GetAllInvoicesFilter = "unpaidoverdue"
	Unbooked      GetAllInvoicesFilter = "unbooked"
)

type GetInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type UpdateInvoiceReq struct {
	Invoice struct {
		AdministrationFee      int    `json:"AdministrationFee"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language         string `json:"Language"`
		NotCompleted     bool   `json:"NotCompleted"`
		OCR              string `json:"OCR"`
		OurReference     string `json:"OurReference"`
		PaymentWay       string `json:"PaymentWay"`
		Phone1           string `json:"Phone1"`
		Phone2           string `json:"Phone2"`
		PriceList        string `json:"PriceList"`
		PrintTemplate    string `json:"PrintTemplate"`
		Project          string `json:"Project"`
		OutboundDate     string `json:"OutboundDate"`
		Remarks          string `json:"Remarks"`
		TermsOfDelivery  string `json:"TermsOfDelivery"`
		TermsOfPayment   string `json:"TermsOfPayment"`
		VATIncluded      bool   `json:"VATIncluded"`
		WayOfDelivery    string `json:"WayOfDelivery"`
		YourOrderNumber  string `json:"YourOrderNumber"`
		YourReference    string `json:"YourReference"`
		ZipCode          string `json:"ZipCode"`
		TaxReductionType string `json:"TaxReductionType"`
	} `json:"Invoice"`
}

type UpdateInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type GetAllInvoicesResp struct {
	Invoices []struct {
		Url                       string `json:"@url"`
		Balance                   int    `json:"Balance"`
		Booked                    bool   `json:"Booked"`
		Cancelled                 bool   `json:"Cancelled"`
		CostCenter                string `json:"CostCenter"`
		Currency                  string `json:"Currency"`
		CurrencyRate              int    `json:"CurrencyRate"`
		CurrencyUnit              int    `json:"CurrencyUnit"`
		CustomerName              string `json:"CustomerName"`
		CustomerNumber            string `json:"CustomerNumber"`
		DocumentNumber            string `json:"DocumentNumber"`
		DueDate                   string `json:"DueDate"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoiceType               string `json:"InvoiceType"`
		NoxFinans                 bool   `json:"NoxFinans"`
		OCR                       string `json:"OCR"`
		VoucherNumber             int    `json:"VoucherNumber"`
		VoucherSeries             string `json:"VoucherSeries"`
		VoucherYear               int    `json:"VoucherYear"`
		WayOfDelivery             string `json:"WayOfDelivery"`
		TermsOfPayment            string `json:"TermsOfPayment"`
		Project                   string `json:"Project"`
		Sent                      bool   `json:"Sent"`
		Total                     int    `json:"Total"`
		FinalPayDate              string `json:"FinalPayDate"`
	} `json:"Invoices"`
}

type CreateInvoiceReq struct {
	Invoice struct {
		AdministrationFee      int    `json:"AdministrationFee"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language         string `json:"Language"`
		NotCompleted     bool   `json:"NotCompleted"`
		OCR              string `json:"OCR"`
		OurReference     string `json:"OurReference"`
		PaymentWay       string `json:"PaymentWay"`
		Phone1           string `json:"Phone1"`
		Phone2           string `json:"Phone2"`
		PriceList        string `json:"PriceList"`
		PrintTemplate    string `json:"PrintTemplate"`
		Project          string `json:"Project"`
		OutboundDate     string `json:"OutboundDate"`
		Remarks          string `json:"Remarks"`
		TermsOfDelivery  string `json:"TermsOfDelivery"`
		TermsOfPayment   string `json:"TermsOfPayment"`
		VATIncluded      bool   `json:"VATIncluded"`
		WayOfDelivery    string `json:"WayOfDelivery"`
		YourOrderNumber  string `json:"YourOrderNumber"`
		YourReference    string `json:"YourReference"`
		ZipCode          string `json:"ZipCode"`
		TaxReductionType string `json:"TaxReductionType"`
	} `json:"Invoice"`
}

type CreateInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type BookKeepInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type CancelInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type CreditInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type SetInvoiceAsSentResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type SetInvoiceAsDoneResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type SendInvoiceAsEmailResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type SendInvoiceAsEPrintResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}

type SendInvoiceAsEInvoiceResp struct {
	Invoice struct {
		Url                    string `json:"@url"`
		UrlTaxReductionList    string `json:"@urlTaxReductionList"`
		AdministrationFee      int    `json:"AdministrationFee"`
		AdministrationFeeVAT   int    `json:"AdministrationFeeVAT"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		Balance                int    `json:"Balance"`
		BasisTaxReduction      int    `json:"BasisTaxReduction"`
		Booked                 bool   `json:"Booked"`
		Cancelled              bool   `json:"Cancelled"`
		City                   string `json:"City"`
		Comments               string `json:"Comments"`
		ContractReference      int    `json:"ContractReference"`
		ContributionPercent    int    `json:"ContributionPercent"`
		ContributionValue      int    `json:"ContributionValue"`
		Country                string `json:"Country"`
		CostCenter             string `json:"CostCenter"`
		Credit                 string `json:"Credit"`
		CreditInvoiceReference string `json:"CreditInvoiceReference"`
		Currency               string `json:"Currency"`
		CurrencyRate           int    `json:"CurrencyRate"`
		CurrencyUnit           int    `json:"CurrencyUnit"`
		CustomerName           string `json:"CustomerName"`
		CustomerNumber         string `json:"CustomerNumber"`
		DeliveryAddress1       string `json:"DeliveryAddress1"`
		DeliveryAddress2       string `json:"DeliveryAddress2"`
		DeliveryCity           string `json:"DeliveryCity"`
		DeliveryCountry        string `json:"DeliveryCountry"`
		DeliveryDate           string `json:"DeliveryDate"`
		DeliveryName           string `json:"DeliveryName"`
		DeliveryZipCode        string `json:"DeliveryZipCode"`
		DocumentNumber         string `json:"DocumentNumber"`
		DueDate                string `json:"DueDate"`
		EDIInformation         struct {
			EDIGlobalLocationNumber         string `json:"EDIGlobalLocationNumber"`
			EDIGlobalLocationNumberDelivery string `json:"EDIGlobalLocationNumberDelivery"`
			EDIInvoiceExtra1                string `json:"EDIInvoiceExtra1"`
			EDIInvoiceExtra2                string `json:"EDIInvoiceExtra2"`
			EDIOurElectronicReference       string `json:"EDIOurElectronicReference"`
			EDIYourElectronicReference      string `json:"EDIYourElectronicReference"`
			EDIStatus                       string `json:"EDIStatus"`
		} `json:"EDIInformation"`
		EmailInformation struct {
			EmailAddressFrom string `json:"EmailAddressFrom"`
			EmailAddressTo   string `json:"EmailAddressTo"`
			EmailAddressCC   string `json:"EmailAddressCC"`
			EmailAddressBCC  string `json:"EmailAddressBCC"`
			EmailSubject     string `json:"EmailSubject"`
			EmailBody        string `json:"EmailBody"`
		} `json:"EmailInformation"`
		EUQuarterlyReport         bool   `json:"EUQuarterlyReport"`
		ExternalInvoiceReference1 string `json:"ExternalInvoiceReference1"`
		ExternalInvoiceReference2 string `json:"ExternalInvoiceReference2"`
		Freight                   int    `json:"Freight"`
		FreightVAT                int    `json:"FreightVAT"`
		Gross                     int    `json:"Gross"`
		HouseWork                 bool   `json:"HouseWork"`
		InvoiceDate               string `json:"InvoiceDate"`
		InvoicePeriodStart        string `json:"InvoicePeriodStart"`
		InvoicePeriodEnd          string `json:"InvoicePeriodEnd"`
		InvoicePeriodReference    string `json:"InvoicePeriodReference"`
		InvoiceRows               []struct {
			AccountNumber          int    `json:"AccountNumber"`
			ArticleNumber          string `json:"ArticleNumber"`
			ContributionPercent    string `json:"ContributionPercent"`
			ContributionValue      string `json:"ContributionValue"`
			CostCenter             string `json:"CostCenter"`
			DeliveredQuantity      string `json:"DeliveredQuantity"`
			Description            string `json:"Description"`
			Discount               int    `json:"Discount"`
			DiscountType           string `json:"DiscountType"`
			HouseWork              bool   `json:"HouseWork"`
			HouseWorkHoursToReport int    `json:"HouseWorkHoursToReport"`
			HouseWorkType          string `json:"HouseWorkType"`
			Price                  int    `json:"Price"`
			PriceExcludingVAT      int    `json:"PriceExcludingVAT"`
			Project                string `json:"Project"`
			RowId                  int    `json:"RowId"`
			StockPointCode         string `json:"StockPointCode"`
			Total                  int    `json:"Total"`
			TotalExcludingVAT      int    `json:"TotalExcludingVAT"`
			Unit                   string `json:"Unit"`
			VAT                    int    `json:"VAT"`
		} `json:"InvoiceRows"`
		InvoiceType string `json:"InvoiceType"`
		Labels      []struct {
			Id int `json:"Id"`
		} `json:"Labels"`
		Language           string `json:"Language"`
		LastRemindDate     string `json:"LastRemindDate"`
		Net                int    `json:"Net"`
		NotCompleted       bool   `json:"NotCompleted"`
		NoxFinans          bool   `json:"NoxFinans"`
		OCR                string `json:"OCR"`
		OfferReference     string `json:"OfferReference"`
		OrderReference     string `json:"OrderReference"`
		OrganisationNumber string `json:"OrganisationNumber"`
		OurReference       string `json:"OurReference"`
		PaymentWay         string `json:"PaymentWay"`
		Phone1             string `json:"Phone1"`
		Phone2             string `json:"Phone2"`
		PriceList          string `json:"PriceList"`
		PrintTemplate      string `json:"PrintTemplate"`
		Project            string `json:"Project"`
		WarehouseReady     bool   `json:"WarehouseReady"`
		OutboundDate       string `json:"OutboundDate"`
		Remarks            string `json:"Remarks"`
		Reminders          int    `json:"Reminders"`
		RoundOff           int    `json:"RoundOff"`
		Sent               bool   `json:"Sent"`
		TaxReduction       int    `json:"TaxReduction"`
		TermsOfDelivery    string `json:"TermsOfDelivery"`
		TermsOfPayment     string `json:"TermsOfPayment"`
		TimeBasisReference int    `json:"TimeBasisReference"`
		Total              int    `json:"Total"`
		TotalToPay         int    `json:"TotalToPay"`
		TotalVAT           int    `json:"TotalVAT"`
		VATIncluded        bool   `json:"VATIncluded"`
		VoucherNumber      int    `json:"VoucherNumber"`
		VoucherSeries      string `json:"VoucherSeries"`
		VoucherYear        int    `json:"VoucherYear"`
		WayOfDelivery      string `json:"WayOfDelivery"`
		YourOrderNumber    string `json:"YourOrderNumber"`
		YourReference      string `json:"YourReference"`
		ZipCode            string `json:"ZipCode"`
		AccountingMethod   string `json:"AccountingMethod"`
		TaxReductionType   string `json:"TaxReductionType"`
		FinalPayDate       string `json:"FinalPayDate"`
	} `json:"Invoice"`
}
