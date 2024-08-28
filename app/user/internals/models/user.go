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
	ID       string
	Name     string
	Email    string
	Username string
	Role     string
	TeamID   int32
	Created  time.Time
}

type UserModelInterface interface {
	Insert(name, email, username string) error
	Delete(email string) error
	Get(username string) (*User, error)
	AssignRole(username, role string) error
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, username string) error {
	query := `
	INSERT INTO user (name, email, username, created_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP);
	`
	_, err := m.DB.Exec(context.Background(), query, name, email, username)
	if err != nil {
		log.Printf("Insert error in pg: %v", err)
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			code, _ := strconv.Atoi(pgError.Code)
			if code == 23505 && strings.Contains(pgError.Message, "users_uc_email") {
				return errors.New("models: duplicate email")
			}
			return err
		}

	}
	return nil
}

func (m *UserModel) Delete(email string) error {
	query := `
	DELETE FROM user
	WHERE email = $1
	`
	_, err := m.DB.Exec(context.Background(), query, email)
	if err != nil {
		log.Println("Delete failed in pg")
		return err
	}
	return nil
}

func (m *UserModel) Get(username string) (*User, error) {
	var user User

	query := `
	SELECT id, name, email, username, role, teamID
	FROM user
	WHERE username = $1;
	`

	err := m.DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.Role, &user.TeamID)
	if err != nil {
		log.Println("Select error in pg")
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) AssignRole(username, role string) error {
	query := `
	UPDATE user SET role = $1
	WHERE username = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, role, username)
	if err != nil {
		log.Println("Update error in pg")
		return err
	}
	return nil
}
