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

var testResp = client.GetMeInformationResp{
	MeInformation: client.MeInformation{
		Id:       "1",
		Name:     "Nick",
		Email:    "nick123@gmail.com",
		SysAdmin: false,
		Locale:   "SE",
	},
}

const debug = false

func TestGetMeInformation(t *testing.T) {
	srv := httptest.NewServer(handlers())
	defer srv.Close()

	c := client.NewClient(client.WithURLOpt(srv.URL))
	meInfo, err := c.GetMeInformation(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	testMeInfo := testResp.MeInformation
	if meInfo.MeInformation != testMeInfo {
		t.Fatal("meInfo != testResp")
	}

	if debug {
		innerMeInfo := meInfo.MeInformation
		fmt.Println(innerMeInfo.Id, "|", testMeInfo.Id)
		fmt.Println(innerMeInfo.Name, "|", testMeInfo.Name)
		fmt.Println(innerMeInfo.Email, "|", testMeInfo.Email)
		fmt.Println(innerMeInfo.SysAdmin, "|", testMeInfo.SysAdmin)
		fmt.Println(innerMeInfo.Locale, "|", testMeInfo.Locale)
	}

}

func handlers() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(testResp)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	})

	return r
}
