package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	fmt.Fprintf(writer, "Content-Type: %s", contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	RequestHader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("powered-by", "sakura  endo dan aruno")
	fmt.Fprint(writer, "ok")
}

func TestResponseHEader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	fmt.Println(response.Header.Get("powered-by"))
}
