package main

import (
	"GoLibraryAPI/api/router"
	"GoLibraryAPI/config"
	"fmt"
	"log"
	"net/http"
)

//  @title          GoLibraryAPI
//  @version        1.0
//  @description    This is a Restful API service for a library application made with Go.

//  @contact.name   Chrystalio
//  @contact.url    https://kristoff-dev.space

//  @host       localhost:3000
//  @basePath   /v1

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
