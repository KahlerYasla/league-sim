package controller

import (
	"football-league-simulation/internal/service"
)

type SimulationController struct {
	service *service.SimulationService
}

func NewSimulationController(service *service.SimulationService) *SimulationController {
	return &SimulationController{service: service}
}

func (c *SimulationController) SimulateWeek(week int) error {
	return c.service.SimulateWeek(week)
}

func (c *SimulationController) SimulateSeason() error {
	return c.service.SimulateSeason()
}
