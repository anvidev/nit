package application

import (
	"encoding/json"
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
	token, err := app.service.Oauth.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := app.service.Oauth.Client(r.Context(), token)
	resp, err := client.Get(oauthUserURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.hxRedirect(w, r, "/")
	return
}
