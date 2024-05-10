package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	http.ListenAndServe(":3000", mux)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!")
}
