package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestLibraryNew(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v", "library")

	resp, err := client.PostForm(route,
		url.Values{"name": {"Library Name"}})

	if err != nil || resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}

func TestLibraryOne(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v", "library", 1)

	resp, err := client.Get(route)
	body, _ := ioutil.ReadAll(resp.Body)
	// t.Log(fmt.Sprintf("%s", body))
	if err != nil || resp.StatusCode != 200 {
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}
