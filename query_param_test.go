package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello, world!")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=sakura", nil)
	recorder := httptest.NewRecorder()
	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
func MultiQueryParameter(writer http.ResponseWriter, request *http.Request) {
	Firtsname := request.URL.Query().Get("first_name")
	Lastname := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", Firtsname, Lastname)
}

func TestMultiQueryParams(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=sakura&last_name=endo", nil)
	recorder := httptest.NewRecorder()
	MultiQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func MultiParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultiParameterValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=sakura&name=endo&name=cantik", nil)
	recorder := httptest.NewRecorder()
	MultiParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
