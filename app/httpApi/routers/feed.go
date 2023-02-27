package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type FeedResource struct{}

// Routes creates a REST router for the todos resource
func (rs FeedResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List) // GET /feed - read a list of feed

	return r
}

func (rs FeedResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog list of stuff.."))
}
