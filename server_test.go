package main

import (
	"net/http"
	"testing"
)

func TestSever(t *testing.T) {
	server := http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
