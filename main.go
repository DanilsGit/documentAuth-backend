package main

import (
	"net/http"

	"github.com/danilsgit/documentAuth-backend/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// create cors
	cors := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
				w.WriteHeader(http.StatusNoContent)
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}

	// Use the cors
	r.Use(cors)

	// Define the routes
	r.Get("/", routes.HomeHandler)
	r.Post("/upload", routes.UploadDocument)
	http.ListenAndServe(":8080", r)
}
