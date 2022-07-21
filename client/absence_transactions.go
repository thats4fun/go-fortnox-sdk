package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-fortnox-sdk/absence_transactions"
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
	filter GetAbsenceTransactionsFilter) ([]*AbsenceTransaction, error) {

	resp := &GetAbsenceTransactionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, absenceTransactionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.AbsenceTransactions, nil
}

// CreateNewAbsenceTransaction does _POST https://api.fortnox.se/3/absencetransactions
//
// req - absence transaction to create
func (c *Client) CreateNewAbsenceTransaction(
	ctx context.Context,
	req *CreateAbsenceTransactionReq) (*AbsenceTransaction, error) {

	resp := &CreateAbsenceTransactionResp{}

	err := c._POST(ctx, absenceTransactionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp.AbsenceTransaction, nil
}

// GetAbsenceTransactionByID does _GET https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
func (c *Client) GetAbsenceTransactionByID(ctx context.Context, id string) (*AbsenceTransaction, error) {
	resp := GetAbsenceTransactionByIDResp{}

	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.AbsenceTransaction, nil
}

// UpdateAbsenceTransactionByID does _PUT https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
//
// updateAbsenceTx - absence transaction to update
func (c *Client) UpdateAbsenceTransactionByID(
	ctx context.Context,
	id string,
	updateAbsenceTx *UpdateAbsenceTransactionReq) (*AbsenceTransaction, error) {

	resp := &UpdateAbsenceTransactionResp{}

	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)

	err := c._PUT(ctx, uri, nil, updateAbsenceTx, resp)
	if err != nil {
		return nil, err
	}

	return resp.AbsenceTransaction, nil
}

// DeleteAbsenceTransactionByID does _DELETE https://api.fortnox.se/3/absencetransactions/{id}
//
// id - identifies the transaction
//
// TODO: update to return deleted tx
func (c *Client) DeleteAbsenceTransactionByID(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", absenceTransactionsURI, id)
	return c._DELETE(ctx, uri)
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
	code AbsenceTransactionStatusCode) (*AbsenceTransaction, error) {

	if !absence_transactions.IsCodeInCodeSet(code.String()) {
		return nil, fmt.Errorf("code should be in code set range: %s", absence_transactions.Codes)
	}

	resp := &GetAbsenceTransactionForEmployeeResp{}

	uri := fmt.Sprintf("%s/%s/%s/%s", absenceTransactionsURI, employeeID, date, code)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.AbsenceTransaction, nil
}

type AbsenceTransactionStatusCode string

func (txCode AbsenceTransactionStatusCode) String() string {
	return string(txCode)
}

type GetAbsenceTransactionsFilter struct {
	EmployeeID string
	Date       string
}

func (p GetAbsenceTransactionsFilter) urlValues() url.Values {
	urlValues := url.Values{}

	if strings.TrimSpace(p.EmployeeID) != "" {
		urlValues[employeeIDParamName] = []string{p.EmployeeID}
	}

	if strings.TrimSpace(p.Date) != "" {
		urlValues[dateParamName] = []string{p.Date}
	}

	return urlValues
}

type GetAbsenceTransactionsResp struct {
	AbsenceTransactions []*AbsenceTransaction `json:"AbsenceTransactions"`
}

type AbsenceTransaction struct {
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
}

type CreateAbsenceTransactionReq struct {
	AbsenceTransaction CreateAbsenceTransaction `json:"AbsenceTransaction"`
}

type CreateAbsenceTransaction struct {
	EmployeeId       string `json:"EmployeeId"`
	CauseCode        string `json:"CauseCode"`
	Date             string `json:"Date"`
	Extent           int    `json:"Extent"`
	Hours            int    `json:"Hours"`
	HolidayEntitling bool   `json:"HolidayEntitling"`
	CostCenter       string `json:"CostCenter"`
	Project          string `json:"Project"`
}

type CreateAbsenceTransactionResp struct {
	AbsenceTransaction *AbsenceTransaction `json:"AbsenceTransaction"`
}

type GetAbsenceTransactionByIDResp CreateAbsenceTransactionResp

type UpdateAbsenceTransactionReq CreateAbsenceTransactionReq

type UpdateAbsenceTransactionResp CreateAbsenceTransactionResp

type GetAbsenceTransactionForEmployeeResp CreateAbsenceTransactionResp
