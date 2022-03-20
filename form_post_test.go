package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// err := request.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }
	firtsname := request.PostFormValue("first_name")
	lastname := request.PostFormValue("last_name")
	// firtname := request.PostForm.Get("first_name")
	// lastname := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firtsname, lastname)

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=sakura&last_name=endo")
	request := httptest.NewRequest("POST", "http://localhost:8080/hello", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
