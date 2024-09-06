package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	NotStarted = "Not Started"
	InProgress = "In Progress"
	Backlog    = "Backlog"
	Complete   = "Complete"
)

type Task struct {
	ID          int32
	Description string
	Creator     int32
	Status      string
	Created     time.Time
	Due         time.Time
}

type ITaskModel interface {
	Insert(id, creator int32, description, status string) error
	Get(id int32) (Task, error)
	GetByCreator(creatorId int32) ([][]Task, error)
	UpdateDescription(id int32, description string) error
	UpdateStatus(id int32, status string) error
	Delete(id int32) error
}

type TaskModel struct {
	DB *pgxpool.Pool
}

func (m *TaskModel) Insert(id, creator int32, description, status string) error {
	return nil
}

func (m *TaskModel) Get(id int32) (Task, error) {
	return Task{}, nil
}

func (m *TaskModel) GetByCreator(creatorId int32) ([][]Task, error) {
	return [][]Task{}, nil
}

func (m *TaskModel) UpdateDescription(id int32, description string) error {
	return nil
}

func (m *TaskModel) UpdateStatus(id int32, status string) error {
	return nil
}

func (m *TaskModel) Delete(id int32) error {
	return nil
}
