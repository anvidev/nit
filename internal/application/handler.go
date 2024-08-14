package application

import (
	"net/http"

	"github.com/anvidev/nit/internal/view/landing"
)

func (app *application) LandingGetPage(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(r, w, landing.ShowLanding())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (app *application) LoginWithFacebook(w http.ResponseWriter, r *http.Request) {
	url := app.service.Oauth.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (app *application) CallbackWithFacebook(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if err := app.service.LoginWithFacebook(w, r, code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app.hxRedirect(w, r, "/")
	return
}
