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

func (s *Service) ListProjectsByID(ID int) ([]Project, error) {
	rows, err := s.DB.Query("SELECT * FROM nit_project WHERE user_id = $1 ORDER BY inserted DESC", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []Project

	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Title, &p.UserID, &p.Inserted); err != nil {
			return projects, nil
		}
		projects = append(projects, p)
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}

	return projects, nil
}

func (s *Service) GetProjectByID(ID string) (Project, error) {
	var proj Project
	err := s.DB.QueryRow("SELECT * FROM nit_project WHERE id = $1", ID).Scan(&proj.ID, &proj.Title, &proj.UserID, &proj.Inserted)
	if err != nil {
		return proj, err
	}
	return proj, nil
}

func (s *Service) DeleteProjectByID(ID string) error {
	_, err := s.DB.Exec("DELETE FROM nit_project WHERE id = $1", ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ListProjects() ([]Project, error) {
	rows, err := s.DB.Query("SELECT * FROM nit_project ORDER BY inserted DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projs []Project
	for rows.Next() {
		var proj Project
		if err := rows.Scan(&proj.ID, &proj.Title, &proj.UserID, &proj.Inserted); err != nil {
			return projs, err
		}
		projs = append(projs, proj)
	}

	if err = rows.Err(); err != nil {
		return projs, err
	}

	return projs, nil
}
