package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-fortnox-sdk/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testRespCompSettings = client.GetCompanySettingResp{
	CompanySettings: client.CompanySettings{
		Address:            "some-address",
		BG:                 "fdfs",
		BIC:                "fdsfs",
		BranchCode:         "fff",
		City:               "NWE",
		ContactFirstName:   "FFF1",
		ContactLastName:    "FFF2",
		Country:            "USA",
		CountryCode:        "123",
		DatabaseNumber:     "#1",
		Domicile:           "gn-2",
		Email:              "testmail.com",
		Fax:                "24254554",
		IBAN:               "-1",
		Name:               "John",
		OrganizationNumber: "1",
		PG:                 "22",
		Phone1:             "-1",
		Phone2:             "-1",
		TaxEnabled:         false,
		VATNumber:          "-1",
		VisitAddress:       "-1",
		VisitCity:          "-1",
		VisitCountry:       "-1",
		VisitCountryCode:   "-1",
		VisitName:          "-1",
		VisitZipCode:       "-1",
		WWW:                "-1",
		ZipCode:            "31234",
	},
}

func TestGetCompanySetting(t *testing.T) {
	srv := httptest.NewServer(compSettingHandlers())
	defer srv.Close()

	testCompSettings := testRespCompSettings.CompanySettings

	c := client.NewClient(client.WithURLOpt(srv.URL))
	compSettings, err := c.GetCompanySettings(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if compSettings.CompanySettings != testCompSettings {
		t.Fatal("compSettings != testCompSettings")
	}
}

func compSettingHandlers() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/settings/company/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(testRespCompSettings)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	})

	return r
}
