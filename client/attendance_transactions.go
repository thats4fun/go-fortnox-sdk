package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	attendanceTransactionsURI = "attendancetransactions"
)

const (
	attendanceTransactionsEmployeeIDParamName = "employeeid"
	attendanceTransactionsDateParamName       = "date"
)

// GetAllAttendanceTransactions does _GET https://api.fortnox.se/3/attendancetransactions
//
// filter - GetAllAttendanceTransactionsFilter
func (c *Client) GetAllAttendanceTransactions(
	ctx context.Context,
	filter *GetAllAttendanceTransactionsFilter) (*GetAllAttendanceTransactionsResp, error) {

	resp := &GetAllAttendanceTransactionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, attendanceTransactionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateAttendanceTransaction does _POST https://api.fortnox.se/3/attendancetransactions
//
// req - attendance transaction to create
func (c *Client) CreateAttendanceTransaction(
	ctx context.Context,
	req *CreateAttendanceTransactionReq) (*CreateAttendanceTransactionResp, error) {

	resp := &CreateAttendanceTransactionResp{}

	err := c._POST(ctx, attendanceTransactionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAttendanceTransaction does _GET https://api.fortnox.se/3/attendancetransactions/{id}
//
// id - identifies the transaction
func (c *Client) GetAttendanceTransaction(ctx context.Context, id string) (*GetAttendanceTransactionResp, error) {
	resp := &GetAttendanceTransactionResp{}

	uri := fmt.Sprintf("%s/%s", attendanceTransactionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateAttendanceTransaction does _PUT https://api.fortnox.se/3/attendancetransactions/{id}
//
// id - identifies the transaction
func (c *Client) UpdateAttendanceTransaction(
	ctx context.Context,
	id string,
	req *UpdateAttendanceTransactionReq) (*UpdateAttendanceTransactionResp, error) {

	resp := &UpdateAttendanceTransactionResp{}

	uri := fmt.Sprintf("%s/%s", attendanceTransactionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllAttendanceTransactionsFilter struct {
	// filter by employee id
	employeeID string
	// filter by date
	date string
}

func (f *GetAllAttendanceTransactionsFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	params := url.Values{}

	if strings.TrimSpace(f.employeeID) != "" {
		params[attendanceTransactionsEmployeeIDParamName] = []string{f.employeeID}
	}

	if strings.TrimSpace(f.employeeID) != "" {
		params[attendanceTransactionsDateParamName] = []string{f.date}
	}

	return params
}

type GetAllAttendanceTransactionsResp struct {
	AttendanceTransactions []struct {
		Url        string `json:"@url"`
		Id         string `json:"id"`
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransactions"`
}

type CreateAttendanceTransactionReq struct {
	AttendanceTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransaction"`
}

type CreateAttendanceTransactionResp struct {
	AttendanceTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransaction"`
}

type GetAttendanceTransactionResp struct {
	AttendanceTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransaction"`
}

type UpdateAttendanceTransactionReq struct {
	AttendanceTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransaction"`
}

type UpdateAttendanceTransactionResp struct {
	AttendanceTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		CauseCode  string `json:"CauseCode"`
		Date       string `json:"Date"`
		Hours      string `json:"Hours"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"AttendanceTransaction"`
}
