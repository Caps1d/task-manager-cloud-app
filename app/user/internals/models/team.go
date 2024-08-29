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
	Manager string
	Members []Member
	Created time.Time
}

const (
	ProjectManager = "PM"
	TeamLead       = "TL"
	Developer      = "Dev"
)

type Member struct {
	ID       int32
	Name     string
	Email    string
	Username string
	Role     string
}

type TeamModelInteface interface {
	Insert(name, manager string) error
	GetTeam(id int32) (*Team, error)
	UpdateName(id int32, name string) error
	UpdateManager(id int32, manager string) error
	UpdateMemberRole(teamID, userID int32, role string) error
	Delete(id int32) error
}

type TeamModel struct {
	DB *pgxpool.Pool
}

func (m *TeamModel) Insert(name, manager string) error {
	query := `
	INSERT INTO teams (name, manager, created_at)
	VALUES ($1, $2, CURRENT_TIMESTAMP);
	`
	_, err := m.DB.Exec(context.Background(), query, name, manager)
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
	query := `
	SELECT t.id, t.name, t.manager, u.id, u.name, u.email, u.username, ut.user_role
	FROM team as t JOIN user_teams as ut ON t.id = ut.team_id JOIN users as u ON ut.user_id = u.id
	WHERE t.id = $1;
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

	var team Team
	members := []Member{}
	teamScanned := false

	for rows.Next() {
		var member Member
		var teamID int32
		var teamName, teamManager string

		err := rows.Scan(teamID, teamName, teamManager, &member.ID, &member.Name, &member.Username, &member.Role)
		if err != nil {
			return nil, err
		}

		if !teamScanned {
			team.ID = teamID
			team.Name = teamName
			team.Manager = teamManager
			teamScanned = true
		}

		members = append(members, member)
	}

	team.Members = members

	return &team, nil
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

func (m *TeamModel) UpdateManager(id int32, manager string) error {
	query := `
	UPDATE users SET manager = $1
	WHERE id = $2;
	`
	_, err := m.DB.Exec(context.Background(), query, manager, id)
	if err != nil {
		log.Printf("Error while updating team manager field in pg: %v", err)
		return err
	}
	return nil
}

func (m *TeamModel) UpdateMemberRole(teamID, userID int32, role string) error {
	query := `
	UPDATE user_teams SET user_role = $1
	WHERE team_id = $2 AND user_id = $3;
	`
	_, err := m.DB.Exec(context.Background(), query, role, teamID, userID)
	if err != nil {
		log.Printf("Error while updating user_role field in pg: %v", err)
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
