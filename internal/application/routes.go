package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) setupRoutes() {
	r := app.mux

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(app.withAuth)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", app.getLanding)
	r.Get("/login", app.getLogin)
	r.Get("/logout", app.getLogout)
	r.Get("/callback", app.getCallback)
	r.Get("/discover", app.getDiscoverProjects)

	r.Group(func(r chi.Router) {
		r.Use(app.mustAuth)
		r.Get("/projects", app.getProjectListByID)
		r.Get("/projects/create", app.getCreateProject)
		r.Post("/projects/create", app.postCreateProject)
		r.Get("/projects/{id}", app.getProjectByID)
		r.Delete("/projects/{id}", app.deleteProjectByID)
	})
}
