package controller

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/service"
)

type LeagueController struct {
	service *service.LeagueService
}

func NewLeagueController(service *service.LeagueService) *LeagueController {
	return &LeagueController{service: service}
}

func (c *LeagueController) CreateLeague(league *domain.League) error {
	return c.service.CreateLeague(league)
}

func (c *LeagueController) GetLeagues() ([]*domain.League, error) {
	return c.service.GetLeagues()
}

func (c *LeagueController) GetLeagueByID(id int) (*domain.League, error) {
	return c.service.GetLeagueByID(id)
}

func (c *LeagueController) UpdateLeague(league *domain.League) error {
	return c.service.UpdateLeague(league)
}
