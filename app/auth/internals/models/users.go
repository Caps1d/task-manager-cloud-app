package models

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int32
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModelInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (int32, error)
	Exists(id int32) (bool, error)
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO users (name, email, password_hash, created_at)
	VALUES($1, $2, $3, CURRENT_TIMESTAMP)
	`
	_, err = m.DB.Exec(context.Background(), query, name, email, string(hashedPassword))
	if err != nil {
		log.Printf("Insert error in pg: %v", err)
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			code, _ := strconv.Atoi(pgError.Code)
			if code == 23505 && strings.Contains(pgError.Message, "users_uc_email") {
				return errors.New("models: duplicate email")
			}
		}
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int32, error) {
	var id int32
	var hashedPassword []byte

	query := `
		SELECT id, password_hash
		FROM users
		WHERE email = $1;
	`
	err := m.DB.QueryRow(context.Background(), query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New("models: invalid credentials")
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), hashedPassword)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, errors.New("models: invalid credentials")
		}
	} else {
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Exists(id int32) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(SELECT true FROM users WHERE id = $1);
	`
	err := m.DB.QueryRow(context.Background(), query, id).Scan(&exists)

	return exists, err
}
