package main

import (
	"GoLibraryAPI/api/router"
	"GoLibraryAPI/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	c := config.New()
	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server on :3000")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed!")
	}
}
