package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const absenceTransactionsURI = "absencetransactions"

const (
	employeeIDParamName = "employeeid"
	dateParamName       = "date"
)

const (
	ASK = "ASK"
	FPE = "FPE"
	FRA = "FRA"
	HAV = "HAV"
	KOV = "KOM"
	MIL = "MIL"
	NAR = "NAR"
	OS1 = "OS1"
	OS2 = "OS2"
	OS3 = "OS3"
	OS4 = "OS4"
	OS5 = "OS5"
	PAP = "PAP"
	PEM = "PEM"
	PER = "PER"
	SEM = "SEM"
	SJK = "SJK"
	SMB = "SMB"
	SVE = "SVE"
	TJL = "TJL"
	UTB = "UTB"
	VAB = "VAB"
)

var AbsenceTransactionsCodes = [...]string{
	"ASK",
	"FPE",
	"FRA",
	"HAV",
	"KOM",
	"MIL",
	"NAR",
	"OS1",
	"OS2",
	"OS3",
	"OS4",
	"OS5",
	"PAP",
	"PEM",
	"PER",
	"SEM",
	"SJK",
	"SMB",
	"SVE",
	"TJL",
	"UTB",
	"VAB",
}

var codeSet = map[string]struct{}{
	"ASK": {},
	"FPE": {},
	"FRA": {},
	"HAV": {},
	"KOM": {},
	"MIL": {},
	"NAR": {},
	"OS1": {},
	"OS2": {},
	"OS3": {},
	"OS4": {},
	"OS5": {},
	"PAP": {},
	"PEM": {},
	"PER": {},
	"SEM": {},
	"SJK": {},
	"SMB": {},
	"SVE": {},
	"TJL": {},
	"UTB": {},
	"VAB": {},
}

func isCodeInCodeSet(code string) bool {
	_, ok := codeSet[code]
	return ok
}

// CauseCodes is usable "CauseCodes"
var CauseCodes = map[string]string{
	"ASK": "Arbetsskada",
	"ATF": "Arbetstidsförkortning",
	"FPE": "Föräldraledig",
	"FRA": "Frånvaro, övrigt (FR1 is used in PAXml)",
	"FR1": "Frånvaro, övrigt (FR1 is used in PAXml)",
	"HAV": "Graviditetspenning",
	"KOM": "Kompledig",
	"MIL": "Militärtjänst (max 60 dagar)",
	"NAR": "Närståendevård (NÄR is used in PAXml)",
	"NÄR": "Närståendevård (NÄR is used in PAXml)",
	"OS1": "Sjuk-OB 1",
	"OS2": "Sjuk-OB 2",
	"OS3": "Sjuk-OB 3",
	"OS4": "Sjuk-OB 4",
	"OS5": "Sjuk-OB 5",
	"PAP": "Pappadagar",
	"PEM": "Permission",
	"PER": "Permitterad",
	"SEM": "Semester",
	"SJK": "Sjukfrånvaro",
	"SMB": "Smittbärare",
	"SVE": "Svenska för invandrare",
	"TJL": "Tjänstledig",
	"UTB": "Facklig utbildning (FAC is used in PAXml)",
	"FAC": "Facklig utbildning (FAC is used in PAXml)",
	"VAB": "Vård av barn",
}

// GetAllAbsenceTransactions does _GET https://api.fortnox.se/3/absencetransactions
//
// filter - may contain employeeID and date
func (c *Client) GetAllAbsenceTransactions(
	ctx context.Context,
	filter *GetAbsenceTransactionsFilter) (*GetAllAbsenceTransactionsResp, error) {

	resp := &GetAllAbsenceTransactionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, absenceTransactionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateNewAbsenceTransaction does _POST https://api.fortnox.se/3/absencetransactions
//
// req - absence transaction to create
func (c *Client) CreateNewAbsenceTransaction(
	ctx context.Context,
	req *CreateNewAbsenceTransactionReq) (*CreateNewAbsenceTransactionResp, error) {

	resp := &CreateNewAbsenceTransactionResp{}

	err := c._POST(ctx, absenceTransactionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAbsenceTransactionByID does _GET https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
func (c *Client) GetAbsenceTransactionByID(ctx context.Context, id string) (*GetAbsenceTransactionByIDResp, error) {
	resp := &GetAbsenceTransactionByIDResp{}

	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateAbsenceTransactionByID does _PUT https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
//
// req - absence transaction to update
func (c *Client) UpdateAbsenceTransactionByID(
	ctx context.Context,
	id string,
	req *UpdateAbsenceTransactionByIDReq) (*UpdateAbsenceTransactionByIDResp, error) {

	resp := &UpdateAbsenceTransactionByIDResp{}

	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteAbsenceTransactionByID does _DELETE https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
func (c *Client) DeleteAbsenceTransactionByID(ctx context.Context, id string) (*DeleteAbsenceTransactionByIDResp, error) {
	resp := &DeleteAbsenceTransactionByIDResp{}

	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)

	err := c._DELETEWithResult(ctx, uri, resp)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAbsenceTransactionForEmployee does _GET https://api.fortnox.se/3/absencetransactions/{EmployeeId}/{Date}/{Code}
//
// employeeID - identifies the employee
//
// date - date of the absence transaction
//
// code - status code of the absence transaction
func (c *Client) GetAbsenceTransactionForEmployee(
	ctx context.Context,
	employeeID, date string,
	code AbsenceTransactionStatusCode) (*GetAbsenceTransactionForEmployeeResp, error) {

	if !isCodeInCodeSet(code.String()) {
		return nil, fmt.Errorf("code should be in code set range: %s", AbsenceTransactionsCodes)
	}

	resp := &GetAbsenceTransactionForEmployeeResp{}

	uri := fmt.Sprintf("%s/%s/%s/%s", absenceTransactionsURI, employeeID, date, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type AbsenceTransactionStatusCode string

func (txCode AbsenceTransactionStatusCode) String() string {
	return string(txCode)
}

type GetAbsenceTransactionsFilter struct {
	EmployeeID string
	Date       string
}

func (p *GetAbsenceTransactionsFilter) urlValues() url.Values {
	if p == nil {
		return nil
	}

	urlValues := url.Values{}

	if strings.TrimSpace(p.EmployeeID) != "" {
		urlValues[employeeIDParamName] = []string{p.EmployeeID}
	}

	if strings.TrimSpace(p.Date) != "" {
		urlValues[dateParamName] = []string{p.Date}
	}

	return urlValues
}

type GetAllAbsenceTransactionsResp struct {
	AbsenceTransactions []struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransactions"`
}

type CreateNewAbsenceTransactionReq struct {
	AbsenceTransaction struct {
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type CreateNewAbsenceTransactionResp struct {
	AbsenceTransaction struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type GetAbsenceTransactionByIDResp struct {
	AbsenceTransaction struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type UpdateAbsenceTransactionByIDReq struct {
	AbsenceTransaction struct {
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type UpdateAbsenceTransactionByIDResp struct {
	AbsenceTransaction struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type DeleteAbsenceTransactionByIDResp struct {
	AbsenceTransaction struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransaction"`
}

type GetAbsenceTransactionForEmployeeResp struct {
	AbsenceTransactions []struct {
		Url              string `json:"@url"`
		Id               string `json:"id"`
		EmployeeId       string `json:"EmployeeId"`
		CauseCode        string `json:"CauseCode"`
		Date             string `json:"Date"`
		Extent           int    `json:"Extent"`
		Hours            int    `json:"Hours"`
		HolidayEntitling bool   `json:"HolidayEntitling"`
		CostCenter       string `json:"CostCenter"`
		Project          string `json:"Project"`
	} `json:"AbsenceTransactions"`
}
