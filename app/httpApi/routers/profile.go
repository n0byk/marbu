package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProfileResource struct{}

// Routes creates a REST router for the todos resource
func (rs ProfileResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)    // GET /profile - read a list of profile
	r.Post("/", rs.Create) // POST /profile - create a new profile and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a profile map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /profile/{id} - read a single profile by :id
		r.Put("/", rs.Update)    // PUT /profile/{id} - update a single profile by :id
		r.Delete("/", rs.Delete) // DELETE /profile/{id} - delete a single profile by :id
	})

	return r
}

func (rs ProfileResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog list of stuff.."))
}

func (rs ProfileResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog create"))
}

func (rs ProfileResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog get"))
}

func (rs ProfileResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog update"))
}

func (rs ProfileResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog delete"))
}
