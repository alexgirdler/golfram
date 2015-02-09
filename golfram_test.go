package golfram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		body, err := ioutil.ReadFile("boston_to_newyork.xml")
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, string(body))
	}))
	defer ts.Close()

	c := NewClient("", ts.URL)
	result, err := c.NewQuery("from boston to newyork")
	fmt.Println(result)

	if err != nil {
		t.Fail()
	}
	if result.Success != "true" {
		t.Fail()
	}
	if result.Numpods != 1 {
		t.Fail()
	}
	if len(result.Assumptions) != 2 {
		t.Fail()
	}
	if len(result.Sources) != 1 {
		t.Fail()
	}
}
