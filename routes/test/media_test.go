package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMediaOne(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v/%v", "media", "info", "566b0b5678b65a57f91fd232")

	resp, err := client.Get(route)
	body, _ := ioutil.ReadAll(resp.Body)
	// t.Log(fmt.Sprintf("%s", body))
	if err != nil || resp.StatusCode != 200 {
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}

func TestMediaAll(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v%v", "media", "info", "?limit=10&skip=0&sort=id")

	resp, err := client.Get(route)
	body, _ := ioutil.ReadAll(resp.Body)
	// t.Log(fmt.Sprintf("%s", body))
	if err != nil || resp.StatusCode != 200 {
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}

func TestMediaPreview(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v", "media", "566b0b5678b65a57f91fd232")

	resp, err := client.Get(route)
	body, _ := ioutil.ReadAll(resp.Body)
	// t.Log(fmt.Sprintf("%s", body))
	if err != nil || resp.StatusCode != 200 {
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}
