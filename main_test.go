package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestGetLiatrio(t *testing.T) {
	req, err := http.NewRequest("GET", "/liatrio", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLiatrio)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := LiatrioResponse{"Automate all the things!", strconv.FormatInt(time.Now().Unix(), 10)}
	var actual LiatrioResponse
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestPostLiatrio(t *testing.T) {
	message := "Hello, Liatrio!"
	reqBody, err := json.Marshal(LiatrioRequest{message})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/liatrio", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostLiatrio)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := LiatrioResponse{message, strconv.FormatInt(time.Now().Unix(), 10)}
	var actual LiatrioResponse
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func TestGetPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPing)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := PingResponse{"liapi", Version, CommitHash, BuildTime, Hostname}
	var actual PingResponse
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

func isEqual(expected, actual LiatrioResponse) bool {
	return expected.Message == actual.Message && expected.Timestamp == actual.Timestamp
}
