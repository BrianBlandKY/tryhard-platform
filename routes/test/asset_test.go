package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	_ "path/filepath"
	"testing"
	"time"
)

var (
	API_URL    = "http://localhost:8080/"
	ASSET_FILE = "/Users/bland/Haste_Store/Test.jpg"
)

func build(url string, filename string) (req *http.Request, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := &bytes.Buffer{}
	writer := multipart.NewWriter(content)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	writer.WriteField("library_id", "1")
	writer.WriteField("size", "1")
	writer.WriteField("date_modified", time.Now().String())

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, _ = http.NewRequest("POST", url, content)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, err
}

func TestAssetFile(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v", "asset", "file")

	req, err := build(route, ASSET_FILE)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}

func TestAssetGetOne(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v", "asset", 1)
	resp, err := client.Get(route)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if err != nil || resp.StatusCode != 200 {
		t.Log("Invalid Status Code", resp.StatusCode)
		t.FailNow()
	}
}

func TestAssetNew(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v/%v", "asset", "new")

	resp, err := client.PostForm(route, url.Values{
		"library_id": {"1"},
		"type":       {".md"},
		"content":    {"random markdown content"},
	})
	if err != nil || resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}

func TestAssetUpdate(t *testing.T) {
	client := &http.Client{}

	route := fmt.Sprintf(API_URL+"%v", "asset")

	resp, err := client.PostForm(route, url.Values{
		"id":      {"1"},
		"type":    {".md"},
		"content": {"random markdown content (UPDATE)"},
	})
	if err != nil || resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Error(err, resp.StatusCode, fmt.Sprintf("%s", body))
		t.FailNow()
	}
	resp.Body.Close()
}
