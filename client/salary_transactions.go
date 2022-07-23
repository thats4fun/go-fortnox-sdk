package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

const (
	salaryTransactionsURI = "salarytransactions"
)

// GetAllSalaryTransactionsForAllEmployees does _GET https://api.fortnox.se/3/salarytransactions
//
// filter - GetAllSalaryTransactionsForAllEmployeesFilter
func (c *Client) GetAllSalaryTransactionsForAllEmployees(
	ctx context.Context,
	filter *GetAllSalaryTransactionsForAllEmployeesFilter) (*GetAllSalaryTransactionsForAllEmployeesResp, error) {

	resp := &GetAllSalaryTransactionsForAllEmployeesResp{}

	params := filter.urlValues()

	err := c._GET(ctx, salaryTransactionsURI, params, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateSalaryTransactionsForEmployee does _POST https://api.fortnox.se/3/salarytransactions
//
// req - salary transaction to create
func (c *Client) CreateSalaryTransactionsForEmployee(
	ctx context.Context, req *CreateSalaryTransactionsForEmployeeReq) (*CreateSalaryTransactionsForEmployeeResp, error) {

	resp := &CreateSalaryTransactionsForEmployeeResp{}

	err := c._POST(ctx, salaryTransactionsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetSalaryTransactionForEmployees does _GET https://api.fortnox.se/3/salarytransactions/{SalaryRow}
//
// salaryRow - identifies the salary transaction
func (c *Client) GetSalaryTransactionForEmployees(
	ctx context.Context, salaryRow int) (*GetSalaryTransactionForEmployeesResp, error) {

	resp := &GetSalaryTransactionForEmployeesResp{}

	uri := fmt.Sprintf("%s/%d", salaryTransactionsURI, salaryRow)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateSalaryTransactionForEmployee does _PUT https://api.fortnox.se/3/salarytransactions/{SalaryRow}
//
// salaryRow - identifies the salary transaction
//
// req - salary transaction to update
func (c *Client) UpdateSalaryTransactionForEmployee(
	ctx context.Context,
	salaryRow int,
	req *UpdateSalaryTransactionForEmployeeReq) (*UpdateSalaryTransactionForEmployeeResp, error) {

	resp := &UpdateSalaryTransactionForEmployeeResp{}

	uri := fmt.Sprintf("%s/%d", salaryTransactionsURI, salaryRow)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteSalaryTransactionForEmployee does _DELETE https://api.fortnox.se/3/salarytransactions/{SalaryRow}
//
// salaryRow - identifies the salary transaction
func (c *Client) DeleteSalaryTransactionForEmployee(
	ctx context.Context, salaryRow int) (*DeleteSalaryTransactionForEmployeeResp, error) {

	resp := &DeleteSalaryTransactionForEmployeeResp{}

	uri := fmt.Sprintf("%s/%d", salaryTransactionsURI, salaryRow)

	err := c._DELETEWithResult(ctx, uri, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllSalaryTransactionsForAllEmployeesFilter struct {
	// filter on employeeId
	EmployeeId string
	// filter on date
	Date string
}

func (f *GetAllSalaryTransactionsForAllEmployeesFilter) urlValues() url.Values {
	if f == nil {
		return nil
	}

	p := url.Values{}

	if strings.TrimSpace(f.EmployeeId) != "" {
		p["employeeId"] = []string{f.EmployeeId}
	}

	if strings.TrimSpace(f.Date) != "" {
		p["date"] = []string{f.Date}
	}

	return p
}

type GetAllSalaryTransactionsForAllEmployeesResp struct {
	SalaryTransactions []struct {
		Url        string `json:"@url"`
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransactions"`
}

type CreateSalaryTransactionsForEmployeeReq struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}

type CreateSalaryTransactionsForEmployeeResp struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}

type GetSalaryTransactionForEmployeesResp struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}

type UpdateSalaryTransactionForEmployeeReq struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}

type UpdateSalaryTransactionForEmployeeResp struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}

type DeleteSalaryTransactionForEmployeeResp struct {
	SalaryTransaction struct {
		EmployeeId string `json:"EmployeeId"`
		SalaryCode string `json:"SalaryCode"`
		SalaryRow  int    `json:"SalaryRow"`
		Date       string `json:"Date"`
		Number     string `json:"Number"`
		Amount     string `json:"Amount"`
		Total      string `json:"Total"`
		Expense    string `json:"Expense"`
		VAT        string `json:"VAT"`
		TextRow    string `json:"TextRow"`
		CostCenter string `json:"CostCenter"`
		Project    string `json:"Project"`
	} `json:"SalaryTransaction"`
}
