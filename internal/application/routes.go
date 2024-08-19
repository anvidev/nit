package application

import (
	"context"
	"net/http"

	"github.com/anvidev/nit/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) setupRoutes() {
	r := app.mux

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.RedirectSlashes)
	r.Use(app.withAuth)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", app.getLanding)
	r.Get("/login", app.getLogin)
	r.Get("/logout", app.getLogout)
	r.Get("/callback", app.getCallback)
	r.Get("/discover", nil)

	r.Group(func(r chi.Router) {
		r.Use(app.mustAuth)
		r.Get("/projects", app.getProjectListByID)
		r.Get("/projects/create", app.getCreateProject)
		r.Post("/projects/create", app.postCreateProject)
		r.Get("/projects/{id}", app.getProjectByID)
		r.Delete("/projects/{id}", app.deleteProjectByID)
	})
}

func (app *application) withAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.service.Store.Get(r, service.CookieKey)
		user, _ := session.Values[service.UserKey].(service.User)

		ctx := context.WithValue(r.Context(), service.UserKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) mustAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := app.getAuthedUser(r.Context())
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
