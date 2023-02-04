package handlers

import (
	"fmt"
	"net/http"

	"github.com/nttn9x/pkg/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	config.App.Session.Put(r.Context(), "message", "Hello from a session!")

	fmt.Fprint(w, "This is home page!")
}

func About(w http.ResponseWriter, r *http.Request) {
	msg := config.App.Session.GetString(r.Context(), "message")

	fmt.Fprint(w, msg)
}
