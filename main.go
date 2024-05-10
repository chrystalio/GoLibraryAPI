package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!")
}
