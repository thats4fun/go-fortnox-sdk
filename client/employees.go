package client

import (
	"context"
	"fmt"
)

const (
	employeesURI = "employees"
)

// GetAllEmployees does _GET https://api.fortnox.se/3/employees/
func (c *Client) GetAllEmployees(ctx context.Context) (*GetAllEmployeesResp, error) {
	resp := &GetAllEmployeesResp{}

	err := c._GET(ctx, employeesURI, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateEmployee does _POST https://api.fortnox.se/3/employees/
//
// req - employee to create
func (c *Client) CreateEmployee(ctx context.Context, req *CreateEmployeeReq) (*CreateEmployeeResp, error) {
	resp := &CreateEmployeeResp{}

	err := c._POST(ctx, employeesURI, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetEmployee does _GET https://api.fortnox.se/3/employees/{EmployeeId}
//
// employeeID - identifies the employee
func (c *Client) GetEmployee(ctx context.Context, employeeID string) (*GetEmployeeResp, error) {
	resp := &GetEmployeeResp{}

	uri := fmt.Sprintf("%s/%s", employeesURI, employeeID)

	err := c._GET(ctx, uri, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateEmployee does _PUT https://api.fortnox.se/3/employees/{EmployeeId}
//
// employeeID - identifies the employee
//
// req - employee to update
func (c *Client) UpdateEmployee(ctx context.Context, employeeID string, req *UpdateEmployeeReq) (*UpdateEmployeeResp, error) {
	resp := &UpdateEmployeeResp{}

	uri := fmt.Sprintf("%s/%s", employeesURI, employeeID)

	err := c._PUT(ctx, uri, nil, req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAllEmployeesResp struct {
	Employees []struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
		Url string `json:"@url"`
	} `json:"Employees"`
}

type CreateEmployeeReq struct {
	Employee struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
	} `json:"Employee"`
}

type CreateEmployeeResp struct {
	Employee struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
	} `json:"Employee"`
}

type GetEmployeeResp struct {
	Employee struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
	} `json:"Employee"`
}

type UpdateEmployeeReq struct {
	Employee struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
	} `json:"Employee"`
}

type UpdateEmployeeResp struct {
	Employee struct {
		EmployeeId             string `json:"EmployeeId"`
		PersonalIdentityNumber string `json:"PersonalIdentityNumber"`
		FirstName              string `json:"FirstName"`
		LastName               string `json:"LastName"`
		FullName               string `json:"FullName"`
		Address1               string `json:"Address1"`
		Address2               string `json:"Address2"`
		PostCode               string `json:"PostCode"`
		City                   string `json:"City"`
		Country                string `json:"Country"`
		Phone1                 string `json:"Phone1"`
		Phone2                 string `json:"Phone2"`
		Email                  string `json:"Email"`
		EmploymentDate         string `json:"EmploymentDate"`
		EmploymentForm         string `json:"EmploymentForm"`
		SalaryForm             string `json:"SalaryForm"`
		JobTitle               string `json:"JobTitle"`
		PersonelType           string `json:"PersonelType"`
		ScheduleId             string `json:"ScheduleId"`
		ForaType               string `json:"ForaType"`
		MonthlySalary          string `json:"MonthlySalary"`
		HourlyPay              string `json:"HourlyPay"`
		TaxAllowance           string `json:"TaxAllowance"`
		TaxTable               string `json:"TaxTable"`
		TaxColumn              int    `json:"TaxColumn"`
		AutoNonRecurringTax    bool   `json:"AutoNonRecurringTax"`
		NonRecurringTax        string `json:"NonRecurringTax"`
		Inactive               bool   `json:"Inactive"`
		ClearingNo             string `json:"ClearingNo"`
		BankAccountNo          string `json:"BankAccountNo"`
		EmployedTo             string `json:"EmployedTo"`
		AverageWeeklyHours     string `json:"AverageWeeklyHours"`
		AverageHourlyWage      string `json:"AverageHourlyWage"`
		DatedWages             []struct {
			EmployeeId    string `json:"EmployeeId"`
			FirstDay      string `json:"FirstDay"`
			MonthlySalary string `json:"MonthlySalary"`
			HourlyPay     string `json:"HourlyPay"`
		} `json:"DatedWages"`
		DatedSchedules []struct {
			EmployeeId string `json:"EmployeeId"`
			FirstDay   string `json:"FirstDay"`
			ScheduleId string `json:"ScheduleId"`
		} `json:"DatedSchedules"`
	} `json:"Employee"`
}
