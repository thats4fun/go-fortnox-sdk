package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/thats4fun/go-fortnox-sdk/absence_transactions"
)

const absenceTransactionsURI = "absencetransactions"

const (
	employeeIDParamName = "employeeid"
	dateParamName       = "date"
)

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

	if !absence_transactions.IsCodeInCodeSet(code.String()) {
		return nil, fmt.Errorf("code should be in code set range: %s", absence_transactions.Codes)
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
