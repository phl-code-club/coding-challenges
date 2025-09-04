package services

import (
	"database/sql"
)

type Team struct {
	ID      int
	Name    string
	Members []string
}

type CreateTeam struct {
	Name    string
	Members []string
}

type TeamService interface {
	CreateTeam(UserInput CreateTeam) (*Team, error)
	GetTeamByID(id int) (*Team, error)
	ListTeams() ([]Team, error)
}

type teamService struct {
	db *sql.DB
}

// CreateTeam implements TeamService.
func (t teamService) CreateTeam(input CreateTeam) (*Team, error) {
	result := t.db.QueryRow("INSERT INTO teams(name) VALUES (?) RETURNING id, name;", input.Name)
	var team Team
	err := result.Scan(&team.ID, &team.Name)
	if err != nil {
		return nil, err
	}
	for name := range input.Members {
		_, err := t.db.Exec("INSERT INTO members(name, team_id) VALUES (?, ?)", name, team.ID)
		if err != nil {
			return nil, err
		}
	}

	team.Members = input.Members
	return &team, nil
}

// GetTeamById implements TeamService.
func (t teamService) GetTeamByID(id int) (*Team, error) {
	result := t.db.QueryRow("SELECT id, name FROM teams WHERE id = ?;", id)
	var team Team
	err := result.Scan(&team.ID, &team.Name)
	if err != nil {
		return nil, err
	}
	memberRows, err := t.db.Query("SELECT name FROM members WHERE team_id = ?", team.ID)
	if err != nil {
		return nil, err
	}
	var members []string
	for memberRows.Next() {
		var member string
		if err := memberRows.Scan(&member); err != nil {
			return nil, err
		}

		members = append(members, member)
	}
	team.Members = members
	return &team, nil
}

// ListTeams implements TeamService.
func (t teamService) ListTeams() ([]Team, error) {
	rows, err := t.db.Query("SELECT id, name FROM teams;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var teams []Team

	for rows.Next() {
		var team Team
		if err := rows.Scan(&team.ID, &team.Name); err != nil {
			return nil, err
		}
		memberRows, err := t.db.Query("SELECT name FROM members WHERE team_id = ?", team.ID)
		if err != nil {
			return nil, err
		}
		members := make([]string, 0)
		var member string
		for memberRows.Next() {
			if err := memberRows.Scan(&member); err != nil {
				return nil, err
			}

			members = append(members, member)
		}
		team.Members = members
		teams = append(teams, team)
	}

	return teams, nil
}

func NewTeamServicxe(db *sql.DB) TeamService {
	return teamService{db}
}
