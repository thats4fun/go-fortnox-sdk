package client

import (
	"context"
	"fmt"
)

const (
	scheduleTimesURI = "scheduletimes"
)

// GetScheduleTime does _GET
//
// employeeID -
//
// date -
func (c *Client) GetScheduleTime(ctx context.Context, employeeID, date string) (*GetScheduleTimeResp, error) {
	resp := &GetScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s", scheduleTimesURI, employeeID, date)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateScheduleTime does _PUT
// employeeID -
// date -
// req -
func (c *Client) UpdateScheduleTime(
	ctx context.Context,
	employeeID, date string,
	req *UpdateScheduleTimeReq) (*UpdateScheduleTimeResp, error) {

	resp := &UpdateScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s", scheduleTimesURI, employeeID, date)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ResetScheduleTime does _PUT
// employeeID -
// date -
func (c *Client) ResetScheduleTime(ctx context.Context, employeeID, date string) (*ResetScheduleTimeResp, error) {
	resp := &ResetScheduleTimeResp{}

	uri := fmt.Sprintf("%s/%s/%s", scheduleTimesURI, employeeID, date)

	err := c._PUT(ctx, uri, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetScheduleTimeResp struct {
	ScheduleTime struct {
		EmployeeId string `json:"EmployeeId"`
		Date       string `json:"Date"`
		ScheduleId string `json:"ScheduleId"`
		Hours      string `json:"Hours"`
		IWH1       string `json:"IWH1"`
		IWH2       string `json:"IWH2"`
		IWH3       string `json:"IWH3"`
		IWH4       string `json:"IWH4"`
		IWH5       string `json:"IWH5"`
	} `json:"ScheduleTime"`
}

type UpdateScheduleTimeReq struct {
	ScheduleTime struct {
		EmployeeId string `json:"EmployeeId"`
		Date       string `json:"Date"`
		ScheduleId string `json:"ScheduleId"`
		Hours      string `json:"Hours"`
		IWH1       string `json:"IWH1"`
		IWH2       string `json:"IWH2"`
		IWH3       string `json:"IWH3"`
		IWH4       string `json:"IWH4"`
		IWH5       string `json:"IWH5"`
	} `json:"ScheduleTime"`
}

type UpdateScheduleTimeResp struct {
	ScheduleTime struct {
		EmployeeId string `json:"EmployeeId"`
		Date       string `json:"Date"`
		ScheduleId string `json:"ScheduleId"`
		Hours      string `json:"Hours"`
		IWH1       string `json:"IWH1"`
		IWH2       string `json:"IWH2"`
		IWH3       string `json:"IWH3"`
		IWH4       string `json:"IWH4"`
		IWH5       string `json:"IWH5"`
	} `json:"ScheduleTime"`
}

type ResetScheduleTimeResp struct {
	ScheduleTime struct {
		EmployeeId string `json:"EmployeeId"`
		Date       string `json:"Date"`
		ScheduleId string `json:"ScheduleId"`
		Hours      string `json:"Hours"`
		IWH1       string `json:"IWH1"`
		IWH2       string `json:"IWH2"`
		IWH3       string `json:"IWH3"`
		IWH4       string `json:"IWH4"`
		IWH5       string `json:"IWH5"`
	} `json:"ScheduleTime"`
}
