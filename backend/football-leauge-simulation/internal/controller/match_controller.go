package controller

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/service"
)

type MatchController struct {
	service *service.MatchService
}

func NewMatchController(service *service.MatchService) *MatchController {
	return &MatchController{service: service}
}

func (c *MatchController) CreateMatch(match *domain.Match) error {
	return c.service.CreateMatch(match)
}

func (c *MatchController) GetMatches() ([]*domain.Match, error) {
	return c.service.GetMatches()
}

func (c *MatchController) GetMatchByID(id int) (*domain.Match, error) {
	return c.service.GetMatchByID(id)
}

func (c *MatchController) UpdateMatch(match *domain.Match) error {
	return c.service.UpdateMatch(match)
}
