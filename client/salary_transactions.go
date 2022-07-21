package client

import "context"

const (
	salaryTransactionsURI = "salarytransactions"
)

func (c *Client) GetAllSalaryTransactionsForAllEmployees(ctx context.Context) {}

func (c *Client) CreateSalaryTransactionsForEmployee(ctx context.Context) {}

func (c *Client) GetSalaryTransactionForEmployees(ctx context.Context) {}

func (c *Client) UpdateSalaryTransactionForEmployee(ctx context.Context) {}

func (c *Client) DeleteSalaryTransactionForEmployee(ctx context.Context) {}
