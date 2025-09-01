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
	name    string
	members []string
}

type TeamService interface {
	CreateTeam(input CreateTeam) (Team, error)
	GetTeamByID(id int) (Team, error)
	ListTeams() ([]Team, error)
}

type teamService struct {
	db *sql.DB
}

// CreateTeam implements TeamService.
func (t teamService) CreateTeam(input CreateTeam) (Team, error) {
	panic("unimplemented")
}

// GetTeamById implements TeamService.
func (t teamService) GetTeamByID(id int) (Team, error) {
	panic("unimplemented")
}

// ListTeams implements TeamService.
func (t teamService) ListTeams() ([]Team, error) {
	panic("unimplemented")
}

func NewTeamServicxe(db *sql.DB) TeamService {
	return teamService{db}
}
