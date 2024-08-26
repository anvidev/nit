package application

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/anvidev/nit/config"
	"github.com/anvidev/nit/internal/service"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
)

const oauthUserURI = "https://graph.facebook.com/v20.0/me?fields=id,name,email"

type application struct {
	config  *config.Config
	logger  *slog.Logger
	mux     *chi.Mux
	service *service.Service
}

func New(cfg *config.Config, logger *slog.Logger, db *sql.DB, s3 *s3.Client) *application {
	return &application{
		config:  cfg,
		logger:  logger,
		mux:     chi.NewRouter(),
		service: service.New(db, s3, cfg),
	}
}

func (app *application) Run() error {
	app.setupRoutes()
  err := http.ListenAndServe(":" + app.config.Addr, app.mux)
	return err
}

func (app *application) renderTemplate(r *http.Request, w http.ResponseWriter, comp templ.Component) error {
	err := comp.Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) hxRedirect(w http.ResponseWriter, r *http.Request, url string) {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

func (app *application) getAuthedUser(ctx context.Context) (*service.User, bool) {
	user, ok := ctx.Value(service.UserKey).(service.User)
	if !ok || len(user.Name) == 0 {
		return nil, false
	}

	return &user, ok
}
