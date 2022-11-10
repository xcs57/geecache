package main

import (
	"net/http"
	"testing"
)

func Test_main(t *testing.T) {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/octet-stream")
		writer.Write([]byte("hello, 47.97.27.101"))
	})
	http.ListenAndServe("localhost:80", nil)

}
