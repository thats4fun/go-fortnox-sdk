package client

import (
	"context"
	"fmt"
)

const (
	expensesURI = "expenses"
)

// GetAllExpenses does _GET https://api.fortnox.se/3/expenses/
func (c *Client) GetAllExpenses(ctx context.Context) (*GetAllExpensesResp, error) {
	resp := &GetAllExpensesResp{}

	err := c._GET(ctx, expensesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateExpense does _POST https://api.fortnox.se/3/expenses/
//
// req - expense to create
func (c *Client) CreateExpense(ctx context.Context, req *CreateExpenseReq) (*CreateExpenseResp, error) {
	resp := &CreateExpenseResp{}

	err := c._POST(ctx, expensesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetExpense does _GET https://api.fortnox.se/3/expenses/{ExpenseCode}
//
// ExpenseCode - expenseCode
func (c *Client) GetExpense(ctx context.Context, expenseCode string) (*GetExpenseResp, error) {
	resp := &GetExpenseResp{}

	uri := fmt.Sprintf("%s/%s", expensesURI, expenseCode)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type Expense struct {
	Code    string `json:"Code"`
	Text    string `json:"Text"`
	Account int    `json:"Account"`
}

type GetAllExpensesResp struct {
	Expenses []struct {
		Code    string `json:"Code"`
		Text    string `json:"Text"`
		Account int    `json:"Account"`
		Url     string `json:"@url"`
	} `json:"Expenses"`
}

type CreateExpenseReq struct {
	Expense struct {
		Code    string `json:"Code"`
		Text    string `json:"Text"`
		Account int    `json:"Account"`
	} `json:"Expense"`
}

type CreateExpenseResp struct {
	Expense struct {
		Code    string `json:"Code"`
		Text    string `json:"Text"`
		Account int    `json:"Account"`
	} `json:"Expense"`
}

type GetExpenseResp struct {
	Expense struct {
		Code    string `json:"Code"`
		Text    string `json:"Text"`
		Account int    `json:"Account"`
	} `json:"Expense"`
}
