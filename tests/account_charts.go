package tests

import (
	"testing"

	"github.com/go-fortnox-sdk/client"
)

var testGetAccountCharts = client.AccountChartResp{
	AccountCharts: []struct {
		Name string `json:"Name"`
	}([]struct {
		Name string `json:"name"`
	}{
		{
			Name: "Acc1",
		},
		{
			Name: "Acc2",
		},
	}),
}

func TestGetAccountCharts(t *testing.T) {

}
