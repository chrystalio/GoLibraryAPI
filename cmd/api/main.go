package main

import (
	"GoLibraryAPI/config"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server on :3000")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed!")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!")
}
