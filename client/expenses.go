package client

import (
	"context"
	"fmt"
)

const (
	expensesURI = "expenses"
)

// GetAllExpenses does _GET https://api.fortnox.se/3/expenses/
func (c *Client) GetAllExpenses(ctx context.Context) ([]Expense, error) {
	resp := &GetAllExpensesResp{}

	err := c._GET(ctx, expensesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp.Expenses, nil
}

// CreateExpense does _POST https://api.fortnox.se/3/expenses/
//
// req - expense to create
func (c *Client) CreateExpense(ctx context.Context, e *Expense) (*Expense, error) {
	req := &CreateExpenseReq{Expense: *e}
	resp := &CreateExpenseResp{}

	err := c._POST(ctx, expensesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Expense, nil
}

// GetExpense does _GET https://api.fortnox.se/3/expenses/{ExpenseCode}
//
// expenseCode - expenseCode
func (c *Client) GetExpense(ctx context.Context, expenseCode string) (*Expense, error) {
	resp := &GetExpenseResp{}

	uri := fmt.Sprintf("%s/%s", expensesURI, expenseCode)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.Expense, nil
}

type Expense struct {
	Code    string `json:"Code,omitempty"`
	Text    string `json:"Text,omitempty"`
	Account int    `json:"Account,omitempty"`
	Url     string `json:"@url,omitempty,omitempty"`
}

type GetAllExpensesResp struct {
	Expenses []Expense `json:"Expenses"`
}

type CreateExpenseReq struct {
	Expense Expense `json:"Expense"`
}

type CreateExpenseResp struct {
	Expense Expense `json:"Expense"`
}

type GetExpenseResp struct {
	Expense Expense `json:"Expense"`
}
