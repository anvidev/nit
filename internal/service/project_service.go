package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Project struct {
	ID          int
	Title       string
	Description string
	Designer    string
	Size        string
	Yarn        string
	Needles     string
	Started     time.Time
	Ended       time.Time
	UserID      int
	Inserted    time.Time
}

type NewProject struct {
	Title       string
	Description string
	Designer    string
	Size        string
	Yarn        string
	Needles     string
	Started     time.Time
	Ended       time.Time
	UserID      int
}

func (s *Service) CreateProject(p *NewProject) (int, error) {
	var id int
	err := s.DB.QueryRow("INSERT INTO nit_project (title, description, designer, yarn, size, needles, started, ended, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", p.Title, p.Description, p.Designer, p.Yarn, p.Size, p.Needles, p.Started, p.Ended, p.UserID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
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
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Designer,
			&p.Size,
			&p.Yarn,
			&p.Needles,
			&p.Started,
			&p.Ended,
			&p.UserID,
			&p.Inserted,
		); err != nil {
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
	err := s.DB.QueryRow("SELECT * FROM nit_project WHERE id = $1", ID).Scan(
		&proj.ID,
		&proj.Title,
		&proj.Description,
		&proj.Designer,
		&proj.Size,
		&proj.Yarn,
		&proj.Needles,
		&proj.Started,
		&proj.Ended,
		&proj.UserID,
		&proj.Inserted,
	)
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
		if err := rows.Scan(
			&proj.ID,
			&proj.Title,
			&proj.Description,
			&proj.Designer,
			&proj.Size,
			&proj.Yarn,
			&proj.Needles,
			&proj.Started,
			&proj.Ended,
			&proj.UserID,
			&proj.Inserted,
		); err != nil {
			return projs, err
		}
		projs = append(projs, proj)
	}

	if err = rows.Err(); err != nil {
		return projs, err
	}

	return projs, nil
}

func (s *Service) UploadImages(bucketName string, projID int, files []*multipart.FileHeader) error {
	errCh := make(chan error, len(files))
	var wg sync.WaitGroup

	for _, f := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			f, err := file.Open()
			if err != nil {
				errCh <- err
				return
			}

			obj := s3.PutObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(fmt.Sprintf("images/%d/%s", projID, file.Filename)),
				Body:   f,
			}

			if _, err := s.S3.PutObject(context.TODO(), &obj); err != nil {
				errCh <- err
				return
			}

		}(f)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
