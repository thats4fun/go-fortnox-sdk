package client

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

const (
	financialYearsURI = "financialyears"
)

// GetAllFinancialYears does _GET https://api.fortnox.se/3/financialyears
//
// date - date to filter on, for example 2020-06-30
func (c *Client) GetAllFinancialYears(ctx context.Context, date *GetAllFinancialYearsFilterDate) (*GetAllFinancialYearsResp, error) {
	resp := &GetAllFinancialYearsResp{}

	var filter url.Values

	if date != nil {
		err := date.validate()
		if err != nil {
			return nil, err
		}
		filter = date.urlValues()
	}

	err := c._GET(ctx, financialYearsURI, filter, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateFinancialYear does _POST https://api.fortnox.se/3/financialyears
//
// financialYear - financial year to create
func (c *Client) CreateFinancialYear(ctx context.Context, financialYear *CreateFinancialYearReq) (*CreateFinancialYearResp, error) {
	resp := &CreateFinancialYearResp{}

	err := c._POST(ctx, financialYearsURI, nil, financialYear, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetFinancialYearByID does _GET https://api.fortnox.se/3/financialyears/{Id}
//
// id - identifies the year
func (c *Client) GetFinancialYearByID(ctx context.Context, id int) (*GetFinancialYearByIDResp, error) {
	resp := &GetFinancialYearByIDResp{}

	uri := fmt.Sprintf("/%s/%d", financialYearsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllFinancialYearsFilterDate struct {
	date string
}

const defaultISOLayout = "2006-01-02"

func (f GetAllFinancialYearsFilterDate) validate() error {
	if _, err := time.Parse(defaultISOLayout, f.date); err != nil {
		return err
	}

	return nil
}

func (f GetAllFinancialYearsFilterDate) urlValues() url.Values {
	urlValues := url.Values{}

	if f.date != "" {
		urlValues["Date"] = []string{f.date}
	}

	return urlValues
}

type GetAllFinancialYearsResp struct {
	FinancialYears []struct {
		Url              string `json:"@url"`
		Id               int    `json:"Id"`
		FromDate         string `json:"FromDate"`
		ToDate           string `json:"ToDate"`
		AccountingMethod string `json:"AccountingMethod"`
		AccountCharts    string `json:"accountCharts"`
	} `json:"FinancialYears"`
}

type CreateFinancialYearReq struct {
	FinancialYear struct {
		Url              string `json:"@url"`
		Id               int    `json:"Id"`
		FromDate         string `json:"FromDate"`
		ToDate           string `json:"ToDate"`
		AccountingMethod string `json:"AccountingMethod"`
		AccountCharts    string `json:"accountCharts"`
	} `json:"FinancialYear"`
}

type CreateFinancialYearResp struct {
	FinancialYear struct {
		Url              string `json:"@url"`
		Id               int    `json:"Id"`
		FromDate         string `json:"FromDate"`
		ToDate           string `json:"ToDate"`
		AccountingMethod string `json:"AccountingMethod"`
		AccountCharts    string `json:"accountCharts"`
	} `json:"FinancialYear"`
}

type GetFinancialYearByIDResp struct {
	FinancialYear struct {
		Url              string `json:"@url"`
		Id               int    `json:"Id"`
		FromDate         string `json:"FromDate"`
		ToDate           string `json:"ToDate"`
		AccountingMethod string `json:"AccountingMethod"`
		AccountCharts    string `json:"accountCharts"`
	} `json:"FinancialYear"`
}
