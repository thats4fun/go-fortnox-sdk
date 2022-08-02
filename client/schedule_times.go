package client

import (
	"context"
	"fmt"
)

const (
	scheduleTimesURI = "scheduletimes"
)

// GetScheduleTime does _GET https://api.fortnox.se/3/scheduletimes/{EmployeeId}/{Date}
//
// employeeID - identifies the employee
//
// date - identifies the date
func (c *Client) GetScheduleTime(ctx context.Context, employeeID, date string) (*ScheduleTime, error) {
	resp := &GetScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s", scheduleTimesURI, employeeID, date)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ScheduleTime, nil
}

// UpdateScheduleTime does _PUT https://api.fortnox.se/3/scheduletimes/{EmployeeId}/{Date}
//
// employeeID - identifies the employee
//
// date - identifies the date
//
// req - schedule time to update
func (c *Client) UpdateScheduleTime(
	ctx context.Context,
	employeeID, date string,
	st *ScheduleTime) (*ScheduleTime, error) {

	req := &UpdateScheduleTimeReq{ScheduleTime: *st}
	resp := &UpdateScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s", scheduleTimesURI, employeeID, date)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ScheduleTime, nil
}

// ResetScheduleTime does _PUT https://api.fortnox.se/3/scheduletimes/{EmployeeId}/{Date}/resetday
//
// employeeID - identifies the employee
//
// date - identifies the date
func (c *Client) ResetScheduleTime(ctx context.Context, employeeID, date string) (*ScheduleTime, error) {
	resp := &ResetScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s/resetday", scheduleTimesURI, employeeID, date)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return &resp.ScheduleTime, nil
}

type ScheduleTime struct {
	EmployeeId string `json:"EmployeeId,omitempty"`
	Date       string `json:"Date,omitempty"`
	ScheduleId string `json:"ScheduleId,omitempty"`
	Hours      string `json:"Hours,omitempty"`
	IWH1       string `json:"IWH1,omitempty"`
	IWH2       string `json:"IWH2,omitempty"`
	IWH3       string `json:"IWH3,omitempty"`
	IWH4       string `json:"IWH4,omitempty"`
	IWH5       string `json:"IWH5,omitempty"`
}

type GetScheduleTimeResp struct {
	ScheduleTime ScheduleTime `json:"ScheduleTime"`
}

type UpdateScheduleTimeReq struct {
	ScheduleTime ScheduleTime `json:"ScheduleTime"`
}

type UpdateScheduleTimeResp struct {
	ScheduleTime ScheduleTime `json:"ScheduleTime"`
}

type ResetScheduleTimeResp struct {
	ScheduleTime ScheduleTime `json:"ScheduleTime"`
}
