package models

import "time"

type User struct {
	ID       string
	Name     string
	Username string
	Email    string
	Created  time.Time
}

type UserModelInterface interface {
	Insert(name, email, password string) error
	Update(name, email, password string) error
	Delete(email string) error
}
