package service

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/repository"
)

type LeagueService struct {
	repo repository.Repository[domain.League]
}

func NewLeagueService(repo repository.Repository[domain.League]) *LeagueService {
	return &LeagueService{repo: repo}
}

func (s *LeagueService) CreateLeague(league *domain.League) error {
	return s.repo.Create(league)
}

func (s *LeagueService) GetLeagues() ([]*domain.League, error) {
	return s.repo.GetAll()
}

func (s *LeagueService) GetLeagueByID(id int) (*domain.League, error) {
	return s.repo.GetByID(id)
}

func (s *LeagueService) UpdateLeague(league *domain.League) error {
	return s.repo.Update(league)
}
