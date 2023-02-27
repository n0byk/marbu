package routers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/n0byk/marbu/config"
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
	ctx := context.Background()
	message, _ := config.App.Storage.FeedList(ctx, "as")
	w.Write([]byte(message))
}
