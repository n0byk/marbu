package httpApi

import (
	"github.com/go-chi/chi/v5"
	"github.com/n0byk/marbu/app/httpApi/routers"
)

func NewHttpApi(r *chi.Mux) *chi.Mux {

	r.Mount("/blog", routers.BlogResource{}.Routes())
	r.Mount("/profile", routers.ProfileResource{}.Routes())
	r.Mount("/feed", routers.FeedResource{}.Routes())

	return r

}
