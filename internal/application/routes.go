package application

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	fileServer(r, "/static", filesDir)

	r.Get("/", app.LandingGetPage)
	r.Get("/login", app.LoginWithFacebook)
	r.Get("/callback", app.CallbackWithFacebook)
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
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
