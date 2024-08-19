package application

import (
	"net/http"

	"github.com/anvidev/nit/internal/service"
	"github.com/anvidev/nit/internal/view/landing"
	"github.com/anvidev/nit/internal/view/projects"
	"github.com/go-chi/chi/v5"
)

func (app *application) getLanding(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(r, w, landing.ShowLanding())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (app *application) getLogin(w http.ResponseWriter, r *http.Request) {
	url := app.service.Oauth.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (app *application) getLogout(w http.ResponseWriter, r *http.Request) {
	if err := app.service.LogoutUser(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (app *application) getCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if err := app.service.LoginWithFacebook(w, r, code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app.hxRedirect(w, r, "/")
	return
}

func (app *application) getCreateProject(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(r, w, projects.ShowCreateProject())
	return
}

func (app *application) postCreateProject(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value(service.UserKey).(service.User)
	title := r.FormValue("title")

	np := service.NewProject{
		Title:  title,
		UserID: user.ID,
	}

	if err := app.service.CreateProject(&np); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.hxRedirect(w, r, "/projects")
	return
}

func (app *application) getProjectListByID(w http.ResponseWriter, r *http.Request) {
	user, _ := app.getAuthedUser(r.Context())
	projs, err := app.service.ListProjectsByID(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	app.renderTemplate(r, w, projects.ListProjects(projs))
	return
}

func (app *application) getProjectByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	proj, err := app.service.GetProjectByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	app.renderTemplate(r, w, projects.ViewProject(proj))
	return
}

func (app *application) deleteProjectByID(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	if err := app.service.DeleteProjectByID(ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.hxRedirect(w, r, "/projects")
	return
}
