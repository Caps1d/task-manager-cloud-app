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

// maybe move elsewhere
const (
	ProjectManager = "PM"
	TeamLead       = "TL"
	Developer      = "Dev"
)

type User struct {
	ID       int32
	Name     string
	Email    string
	Username string
	Role     string
	TeamID   int32
	Created  time.Time
}

type UserModelInterface interface {
	Insert(name, email, username string) error
	Get(id int32) (*User, error)
	UpdateEmail(id int32, email string) error
	UpdateRole(id int32, role string) error
	UpdateTeamID(id int32, teamID int32) error
	Delete(id int32) error
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

func (m *UserModel) Delete(id int32) error {
	query := `
	DELETE FROM users
	WHERE id = $1
	`
	_, err := m.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Error in User Delete method in pg: %v", err)
		return err
	}
	return nil
}

func (m *UserModel) Get(id int32) (*User, error) {
	var user User

	query := `
	SELECT id, name, email, username, role, teamID
	FROM users
	WHERE id = $1;
	`

	err := m.DB.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Role, &user.TeamID)
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

func (m *UserModel) UpdateRole(id int32, role string) error {
	query := `
	UPDATE users SET role = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, role, id)
	if err != nil {
		log.Printf("Error while updating role in pg: %v", err)
		return err
	}
	return nil
}

func (m *UserModel) UpdateTeamID(id int32, teamID int32) error {
	query := `
	UPDATE users SET teamID = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, teamID, id)
	if err != nil {
		log.Printf("Error while updating teamID in pg: %v", err)
		return err
	}
	return nil
}

func (m *UserModel) Exist(id int32) (bool, error) {
	var exists bool

	query := `
	SELECT EXISTS(SELECT true FROM users WHERE id = $1);
	`
	err := m.DB.QueryRow(context.Background(), query, id).Scan(exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
