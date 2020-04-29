package main

import (
	"log"
	"net/http"
	"time"

	"github.com/matheusfm/password-validator/api/routers"
)

func main() {
	srv := &http.Server{
		Handler:      routers.Routers(),
		Addr:         ":8000",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	log.Printf("password-validator is ready to handle requests at %s", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
