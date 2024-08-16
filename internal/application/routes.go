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
	r.Use(app.withAuth)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", app.LandingGetPage)
	r.Get("/login", app.LoginWithFacebook)
	r.Get("/callback", app.CallbackWithFacebook)

	r.Group(func(r chi.Router) {
		r.Get("/create", app.createProjectPage)
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
