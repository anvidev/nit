package service

import "time"

type Project struct {
	ID       int
	Title    string
	UserID   int
	Inserted time.Time
}

type NewProject struct {
	Title  string
	UserID int
}

func (s *Service) CreateProject(p *NewProject) error {
	_, err := s.DB.Exec("INSERT INTO nit_project (title, user_id) VALUES ($1, $2)", p.Title, p.UserID)
	if err != nil {
		return err
	}
	return nil
}
