package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/SudoAaron/golang-auth0-template/platform/authenticator"
	"github.com/SudoAaron/golang-auth0-template/platform/router"
)

func init() {
	// Register the type for the session
	gob.Register(map[string]interface{}{})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
