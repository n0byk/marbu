package httpApi

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/n0byk/marbu/app/httpApi/routers"
	"github.com/n0byk/marbu/config"
)

func New() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	mountRoutes(r)
	log.Println("Http api init.")
	if err := http.ListenAndServe(config.Env.APP_PORT, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}

}

func mountRoutes(r *chi.Mux) {
	r.Mount("/blog", routers.BlogResource{}.Routes())
	r.Mount("/profile", routers.ProfileResource{}.Routes())
	r.Mount("/feed", routers.FeedResource{}.Routes())
}
