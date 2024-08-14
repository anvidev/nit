package service

import (
	"database/sql"
	"encoding/gob"

	"github.com/anvidev/nit/config"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Service struct {
	DB    *sql.DB
	Oauth *oauth2.Config
	Store *sessions.CookieStore
}

func New(db *sql.DB, cfg *config.Config) *Service {
	gob.Register(User{})
	return &Service{
		DB: db,
		Oauth: &oauth2.Config{
			ClientID:     cfg.FacebookClientID,
			ClientSecret: cfg.FacebookClientSecret,
			RedirectURL:  "http://localhost:3000/callback",
			Endpoint:     facebook.Endpoint,
		},
		Store: sessions.NewCookieStore([]byte(cfg.SessionSecret)),
	}
}
