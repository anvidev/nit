package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const (
	oauthUserURI = "https://graph.facebook.com/v20.0/me?fields=id,name,email"
	CookieKey    = "nit_session"
	UserKey      = "user"
)

type User struct {
	ID         int
	FacebookID int64
	Name       string
	Email      string
	CreatedAt  time.Time
}

func (s *Service) LoginWithFacebook(w http.ResponseWriter, r *http.Request, code string) error {
	token, err := s.Oauth.Exchange(r.Context(), code)
	if err != nil {
		return err
	}

	client := s.Oauth.Client(r.Context(), token)
	resp, err := client.Get(oauthUserURI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return err
	}

	_, err = s.DB.Exec("INSERT INTO nit_user (facebook_id, name, email) VALUES ($1, $2, $3)", userInfo.ID, userInfo.Name, userInfo.Email)
	if err != nil && !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
		return err
	}

	var user User
	err = s.DB.QueryRow("SELECT * FROM nit_user WHERE facebook_id = $1", userInfo.ID).Scan(&user.ID, &user.FacebookID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return err
	}

	session, err := s.Store.Get(r, CookieKey)
	session.Values[UserKey] = user
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) LogoutUser(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.Store.Get(r, CookieKey)
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
