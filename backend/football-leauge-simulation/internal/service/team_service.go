package service

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/repository"
)

type TeamService struct {
	repo repository.Repository[domain.Team]
}

func NewTeamService(repo repository.Repository[domain.Team]) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateTeam(team *domain.Team) error {
	return s.repo.Create(team)
}

func (s *TeamService) GetTeams() ([]*domain.Team, error) {
	return s.repo.GetAll()
}

func (s *TeamService) GetTeamByID(id int) (*domain.Team, error) {
	return s.repo.GetByID(id)
}

func (s *TeamService) UpdateTeam(team *domain.Team) error {
	return s.repo.Update(team)
}
