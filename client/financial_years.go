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
func (c *Client) GetAllFinancialYears(ctx context.Context, date *GetAllFinancialYearsFilterDate) ([]FinancialYear, error) {
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

	return resp.FinancialYears, nil
}

// CreateFinancialYear does _POST https://api.fortnox.se/3/financialyears
//
// financialYear - financial year to create
func (c *Client) CreateFinancialYear(ctx context.Context, fy *FinancialYear) (*FinancialYear, error) {
	req := &CreateFinancialYearReq{FinancialYear: *fy}
	resp := &CreateFinancialYearResp{}

	err := c._POST(ctx, financialYearsURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.FinancialYear, nil
}

// GetFinancialYearByID does _GET https://api.fortnox.se/3/financialyears/{Id}
//
// id - identifies the year
func (c *Client) GetFinancialYearByID(ctx context.Context, id int) (*FinancialYear, error) {
	resp := &GetFinancialYearByIDResp{}

	uri := fmt.Sprintf("/%s/%d", financialYearsURI, id)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.FinancialYear, nil
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

type FinancialYear struct {
	Url              string `json:"@url,omitempty"`
	Id               int    `json:"Id,omitempty"`
	FromDate         string `json:"FromDate,omitempty"`
	ToDate           string `json:"ToDate,omitempty"`
	AccountingMethod string `json:"AccountingMethod,omitempty"`
	AccountCharts    string `json:"accountCharts,omitempty"`
}

type GetAllFinancialYearsResp struct {
	FinancialYears []FinancialYear `json:"FinancialYears"`
}

type CreateFinancialYearReq struct {
	FinancialYear FinancialYear `json:"FinancialYear"`
}

type CreateFinancialYearResp struct {
	FinancialYear FinancialYear `json:"FinancialYear"`
}

type GetFinancialYearByIDResp struct {
	FinancialYear FinancialYear `json:"FinancialYear"`
}
