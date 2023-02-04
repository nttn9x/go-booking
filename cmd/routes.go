package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nttn9x/pkg/config"
	"github.com/nttn9x/pkg/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// Wrap handlers with the LoadAndSave() middleware.
	mux.Use(func(h http.Handler) http.Handler {
		return config.App.Session.LoadAndSave(h)
	})

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}
