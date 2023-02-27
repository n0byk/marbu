package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BlogResource struct{}

// Routes creates a REST router for the todos resource
func (rs BlogResource) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Use() // some middleware..

	r.Get("/", rs.List)    // GET /blog - read a list of blog
	r.Post("/", rs.Create) // POST /blog - create a new blog and persist it
	r.Put("/", rs.Delete)

	r.Route("/{id}", func(r chi.Router) {
		// r.Use(rs.TodoCtx) // lets have a blog map, and lets actually load/manipulate
		r.Get("/", rs.Get)       // GET /blog/{id} - read a single blog by :id
		r.Put("/", rs.Update)    // PUT /blog/{id} - update a single blog by :id
		r.Delete("/", rs.Delete) // DELETE /blog/{id} - delete a single blog by :id
	})

	return r
}

func (rs BlogResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog list of stuff.."))
}

func (rs BlogResource) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog create"))
}

func (rs BlogResource) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog get"))
}

func (rs BlogResource) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog update"))
}

func (rs BlogResource) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog delete"))
}
