package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/csrf"
	"github.com/joho/godotenv"
)

func main() {
	// Load env vars.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Add CSRF protection. In prod Secure is set to true.
	CSRF := csrf.Protect([]byte(os.Getenv("CSRF_KEY")), csrf.Secure(false))

	// Add all routes.
	r := router()

	// Configure the server.
	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        CSRF(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// Start the server.
	log.Printf("starting server at port %s ...", os.Getenv("PORT"))
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("error running server:", err)
	}
}
