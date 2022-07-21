package tests

import (
	"github.com/go-fortnox-sdk/client"
	"testing"
)

var testGetAllAbsenceTransactionsResp = client.GetAbsenceTransactionsResp{
	AbsenceTransactions: []*client.AbsenceTransaction{
		{
			Url:              "",
			Id:               "",
			EmployeeId:       "",
			CauseCode:        "",
			Date:             "",
			Extent:           0,
			Hours:            0,
			HolidayEntitling: false,
			CostCenter:       "",
			Project:          "",
		},
		{
			Url:              "",
			Id:               "",
			EmployeeId:       "",
			CauseCode:        "",
			Date:             "",
			Extent:           0,
			Hours:            0,
			HolidayEntitling: false,
			CostCenter:       "",
			Project:          "",
		},
	},
}

func TestGetAllAbsenceTransactions(t *testing.T) {

}

func TestCreateNewAbsenceTransaction(t *testing.T) {

}

func TestGetAbsenceTransactionByID(t *testing.T) {

}

func TestUpdateAbsenceTransactionByID(t *testing.T) {

}

func TestDeleteAbsenceTransactionByID(t *testing.T) {

}

func TestGetAbsenceTransactionForEmployee(t *testing.T) {

}
