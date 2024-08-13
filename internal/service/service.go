package service

import (
	"database/sql"

	"github.com/anvidev/nit/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Service struct {
	DB    *sql.DB
	Oauth *oauth2.Config
}

func New(db *sql.DB, cfg *config.Config) *Service {
	return &Service{
		DB: db,
		Oauth: &oauth2.Config{
			ClientID:     cfg.FacebookClientID,
			ClientSecret: cfg.FacebookClientSecret,
			RedirectURL:  "http://localhost:3000/callback",
			Endpoint:     facebook.Endpoint,
		},
	}
}
