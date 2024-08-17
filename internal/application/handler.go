package application

import (
	"net/http"

	"github.com/anvidev/nit/internal/view/landing"
)

func (app *application) handleLandingPage(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(r, w, landing.ShowLanding())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (app *application) handleLogin(w http.ResponseWriter, r *http.Request) {
	url := app.service.Oauth.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (app *application) handleLogout(w http.ResponseWriter, r *http.Request) {
	if err := app.service.LogoutUser(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (app *application) handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if err := app.service.LoginWithFacebook(w, r, code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app.hxRedirect(w, r, "/")
	return
}

func (app *application) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("create page"))
	return
}
