package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Project struct {
	ID          int32
	Description string
	Manager     int32
	Created     time.Time
}

type IProjectModel interface {
	Insert(id, manager int32, description, status string) error
	Get(id int32) (Project, error)
	GetByUser(userId int32) ([][]Project, error)
	UpdateDescription(id int32, description string) error
	Delete(id int32) error
}

type ProjectModel struct {
	DB *pgxpool.Pool
}

func (m *ProjectModel) Insert(id, manager int32, description, status string) error {
	return nil
}

func (m *ProjectModel) Get(id int32) (Project, error) {
	return Project{}, nil
}

func (m *ProjectModel) GetByUser(userId int32) ([][]Project, error) {
	return [][]Project{}, nil
}

func (m *ProjectModel) UpdateDescription(id int32, description string) error {
	return nil
}

func (m *ProjectModel) Delete(id int32) error {
	return nil
}
