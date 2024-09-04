package models

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int32
	Name     string
	Email    string
	Username string
	Created  time.Time
}

type UserModelInterface interface {
	Insert(name, email, username string) error
	Get(id int32) (*User, error)
	UpdateEmail(id int32, email string) error
	Exist(id int32) (bool, error)
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, username string) error {
	query := `
	INSERT INTO users (name, email, username, created_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP);
	`
	_, err := m.DB.Exec(context.Background(), query, name, email, username)
	if err != nil {
		log.Printf("Insert error in pg: %v", err)
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			code, _ := strconv.Atoi(pgError.Code)
			if code == 23505 && strings.Contains(pgError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
			return err
		}

	}
	return nil
}

func (m *UserModel) Get(id int32) (*User, error) {
	var user User

	query := `
	SELECT id, name, email, username
	FROM users
	WHERE id = $1;
	`

	err := m.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Username)
	if err != nil {
		log.Printf("Error in User Get method in pg: %v", err)
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) UpdateEmail(id int32, email string) error {
	query := `
	UPDATE users SET email = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, email, id)
	if err != nil {
		log.Printf("Error while updating email in pg: %v", err)
		return err
	}
	return nil
}

func (m *UserModel) Exist(id int32) (bool, error) {
	var exists bool

	query := `
	SELECT EXISTS(SELECT true FROM users WHERE id = $1);
	`
	err := m.DB.QueryRow(context.Background(), query, id).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
