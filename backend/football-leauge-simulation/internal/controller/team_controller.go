package controller

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/service"
)

type TeamController struct {
	service *service.TeamService
}

func NewTeamController(service *service.TeamService) *TeamController {
	return &TeamController{service: service}
}

func (c *TeamController) CreateTeam(team *domain.Team) error {
	return c.service.CreateTeam(team)
}

func (c *TeamController) GetTeams() ([]*domain.Team, error) {
	return c.service.GetTeams()
}

func (c *TeamController) GetTeamByID(id int) (*domain.Team, error) {
	return c.service.GetTeamByID(id)
}

func (c *TeamController) UpdateTeam(team *domain.Team) error {
	return c.service.UpdateTeam(team)
}
