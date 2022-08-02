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

var _AttendanceTransactionsCauseCodes = map[string]string{
	"ARB": "Timlön",
	"BE2": "Beredskapstid 2",
	"BER": "Beredskapstid (BE1 is used in PAXml)",
	"BE1": "Beredskapstid (BE1 is used in PAXml)",
	"FLX": "Flextid +/-",
	"HLG": "Helglön",
	"JO2": "Jourtid 2 (JR2 is used in PAXml)",
	"JR2": "Jourtid 2 (JR2 is used in PAXml)",
	"JOR": "Jourtid (JR1 is used in PAXml)",
	"JR1": "Jourtid (JR1 is used in PAXml)",
	"MER": "Mertid",
	"OB1": "OB-ersättning 1",
	"OB2": "OB-ersättning 2",
	"OB3": "OB-ersättning 3",
	"OB4": "OB-ersättning 4",
	"OB5": "OB-ersättning 5",
	"OK0": "Extratid \u2013 Komptid (NV9 is used in PAXml)",
	"NV9": "Extratid – Komptid (NV9 is used in PAXml)",
	"OK1": "Övertid 1 \u2013 Komptid (ÖK1 is used in PAXml)",
	"ÖK1": "Övertid 1 – Komptid (ÖK1 is used in PAXml)",
	"OK2": "Övertid 2 \u2013 Komptid (ÖK2 is used in PAXml)",
	"ÖK2": "Övertid 2 – Komptid (ÖK2 is used in PAXml)",
	"OK3": "Övertid 3 \u2013 Komptid (ÖK3 is used in PAXml)",
	"ÖK3": "Övertid 3 – Komptid (ÖK3 is used in PAXml)",
	"OK4": "Övertid 4 \u2013 Komptid (ÖK4 is used in PAXml)",
	"ÖK4": "Övertid 4 – Komptid (ÖK4 is used in PAXml)",
	"OK5": "Övertid 5 \u2013 Komptid (ÖK5 is used in PAXml)",
	"ÖK5": "Övertid 5 – Komptid (ÖK5 is used in PAXml)",
	"OT1": "Övertid 1 \u2013 Betalning (ÖT1 is used in PAXml)",
	"ÖT1": "Övertid 1 – Betalning (ÖT1 is used in PAXml)",
	"OT2": "Övertid 2 \u2013 Betalning (ÖT2 is used in PAXml)",
	"ÖT2": "Övertid 2 – Betalning (ÖT2 is used in PAXml)",
	"OT3": "Övertid 3 \u2013 Betalning (ÖT3 is used in PAXml)",
	"ÖT3": "Övertid 3 – Betalning (ÖT3 is used in PAXml)",
	"OT4": "Övertid 4 \u2013 Betalning (ÖT4 is used in PAXml)",
	"ÖT4": "Övertid 4 – Betalning (ÖT4 is used in PAXml)",
	"OT5": "Övertid 5 \u2013 Betalning (ÖT5 is used in PAXml)",
	"ÖT5": "Övertid 5 – Betalning (ÖT5 is used in PAXml)",
	"RES": "Restid (RE1 is used in PAXml)",
	"RE1": "Restid (RE1 is used in PAXml)",
	"TID": "Arbetstid",
}

// GetAllAttendanceTransactions does _GET https://api.fortnox.se/3/attendancetransactions
//
// filter - GetAllAttendanceTransactionsFilter
func (c *Client) GetAllAttendanceTransactions(
	ctx context.Context,
	filter *GetAllAttendanceTransactionsFilter) ([]AttendanceTransaction, error) {

	resp := &GetAllAttendanceTransactionsResp{}

	params := filter.urlValues()

	err := c._GET(ctx, attendanceTransactionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp.AttendanceTransactions, nil
}

// CreateAttendanceTransaction does _POST https://api.fortnox.se/3/attendancetransactions
//
// at - attendance transaction to create
func (c *Client) CreateAttendanceTransaction(
	ctx context.Context,
	at *AttendanceTransaction) (*AttendanceTransaction, error) {

	req := CreateAttendanceTransactionReq{AttendanceTransaction: *at}
	resp := &CreateAttendanceTransactionResp{}

	err := c._POST(ctx, attendanceTransactionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.AttendanceTransaction, nil
}

// GetAttendanceTransaction does _GET https://api.fortnox.se/3/attendancetransactions/{id}
//
// id - identifies the transaction
func (c *Client) GetAttendanceTransaction(ctx context.Context, id string) (*AttendanceTransaction, error) {
	resp := &GetAttendanceTransactionResp{}

	uri := fmt.Sprintf("%s/%s", attendanceTransactionsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.AttendanceTransaction, nil
}

// UpdateAttendanceTransaction does _PUT https://api.fortnox.se/3/attendancetransactions/{id}
//
// id - identifies the transaction
//
// at - attendance transaction to update
func (c *Client) UpdateAttendanceTransaction(
	ctx context.Context,
	id string,
	at *AttendanceTransaction) (*AttendanceTransaction, error) {

	req := &UpdateAttendanceTransactionReq{AttendanceTransaction: *at}
	resp := &UpdateAttendanceTransactionResp{}

	uri := fmt.Sprintf("%s/%s", attendanceTransactionsURI, id)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.AttendanceTransaction, nil
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

type AttendanceTransaction struct {
	Url        string `json:"@url,omitempty,omitempty"`
	Id         string `json:"id,omitempty,omitempty"`
	EmployeeId string `json:"EmployeeId,omitempty"`
	CauseCode  string `json:"CauseCode,omitempty"`
	Date       string `json:"Date,omitempty"`
	Hours      string `json:"Hours,omitempty"`
	CostCenter string `json:"CostCenter,omitempty"`
	Project    string `json:"Project,omitempty"`
}

type GetAllAttendanceTransactionsResp struct {
	AttendanceTransactions []AttendanceTransaction `json:"AttendanceTransactions"`
}

type CreateAttendanceTransactionReq struct {
	AttendanceTransaction AttendanceTransaction `json:"AttendanceTransaction"`
}

type CreateAttendanceTransactionResp struct {
	AttendanceTransaction AttendanceTransaction `json:"AttendanceTransaction"`
}

type GetAttendanceTransactionResp struct {
	AttendanceTransaction AttendanceTransaction `json:"AttendanceTransaction"`
}

type UpdateAttendanceTransactionReq struct {
	AttendanceTransaction AttendanceTransaction `json:"AttendanceTransaction"`
}

type UpdateAttendanceTransactionResp struct {
	AttendanceTransaction AttendanceTransaction `json:"AttendanceTransaction"`
}
