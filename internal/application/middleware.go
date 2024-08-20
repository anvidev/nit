package application

import (
	"context"
	"net/http"
	"strings"

	"github.com/anvidev/nit/internal/service"
)

// withAuth is a middleware that gets the user data from the current cookiestore session,
// puts it into a [context.Context] from the parent context, and returns shallow copy of
// the parent context with the user data.
//
// This middleware does not care if a user is logged in or not.
func (app *application) withAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/static") {
			h.ServeHTTP(w, r)
			return
		}
		session, _ := app.service.Store.Get(r, service.CookieKey)
		user, _ := session.Values[service.UserKey].(service.User)

		ctx := context.WithValue(r.Context(), service.UserKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// mustAuth is a middleware that gets the current user data from the request context.
// If no user data is found, it redirects the client to homepage, else it continues.
func (app *application) mustAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/static") {
			h.ServeHTTP(w, r)
			return
		}
		_, ok := app.getAuthedUser(r.Context())
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
