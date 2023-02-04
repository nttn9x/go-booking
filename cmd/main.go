package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nttn9x/pkg/config"
)

var portNumber = ":8080"
var sessionManager *scs.SessionManager

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	config.Init(sessionManager)

	fmt.Printf("Starting application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
