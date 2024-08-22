package service

import (
	"database/sql"
	"encoding/gob"

	"github.com/anvidev/nit/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Service struct {
	DB    *sql.DB
	Oauth *oauth2.Config
	Store *sessions.CookieStore
	S3    *s3.Client
}

func New(db *sql.DB, s3 *s3.Client, cfg *config.Config) *Service {
	gob.Register(User{})
	return &Service{
		DB: db,
		Oauth: &oauth2.Config{
			ClientID:     cfg.FacebookClientID,
			ClientSecret: cfg.FacebookClientSecret,
			RedirectURL:  cfg.FacebookCallbackURL,
			Endpoint:     facebook.Endpoint,
		},
		Store: sessions.NewCookieStore([]byte(cfg.SessionSecret)),
		S3:    s3,
	}
}
