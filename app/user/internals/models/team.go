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
)

type Team struct {
	ID      int32
	Name    string
	About   string
	Created time.Time
}

type TeamModelInteface interface {
	Insert(name, about string) error
	GetTeam(id int32) (*Team, error)
	GetTeamMembers(teamID int32) ([]*User, error)
	UpdateName(id int32, name string) error
	UpdateAbout(id int32, about string) error
	Delete(id int32) error
}

type TeamModel struct {
	DB *pgxpool.Pool
}

func (m *TeamModel) Insert(name, about string) error {
	query := `
	INSERT INTO teams (name, about, created_at)
	VALUES ($1, $2, CURRENT_TIMESTAMP);
	`
	_, err := m.DB.Exec(context.Background(), query, name, about)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			code, _ := strconv.Atoi(pgErr.Code)
			if code == 23505 && strings.Contains(pgErr.Message, "teams_uc_name") {
				return ErrDuplicateName
			}
			return err
		}
	}
	return nil
}

func (m *TeamModel) GetTeam(id int32) (*Team, error) {
	var team Team

	query := `
	SELECT id, name, about, created_at FROM teams
	WHERE id = $1;
	`

	err := m.DB.QueryRow(context.Background(), query, id).Scan(&team.ID, &team.Name, &team.About, &team.Created)
	if errors.Is(err, pgx.ErrNoRows) || team.ID != id {
		return nil, ErrNoRecord
	}

	return &team, nil
}

func (m *TeamModel) GetTeamMembers(id int32) ([]*User, error) {
	query := `
	SELECT id, name, username, role, teamID
	FROM users
	WHERE teamID = $1;
	`
	rows, err := m.DB.Query(context.Background(), query, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Role, &user.TeamID)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (m *TeamModel) UpdateName(id int32, name string) error {
	query := `
	UPDATE users SET name = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, name, id)
	if err != nil {
		log.Printf("Error while updating team name in pg: %v", err)
		return err
	}
	return nil
}

func (m *TeamModel) UpdateAbout(id int32, about string) error {
	query := `
	UPDATE users SET about = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, about, id)
	if err != nil {
		log.Printf("Error while updating team about field in pg: %v", err)
		return err
	}
	return nil
}

func (m *TeamModel) Delete(id int32) error {
	query := `
	DELETE FROM teams WHERE id = $1;
	`
	_, err := m.DB.Exec(context.Background(), query, id)
	if err != nil {
		log.Printf("Error in Team Delete method in pg: %v", err)
		return err
	}
	return nil
}
