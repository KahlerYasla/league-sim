package service

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/repository"
)

type MatchService struct {
	repo repository.Repository[domain.Match]
}

func NewMatchService(repo repository.Repository[domain.Match]) *MatchService {
	return &MatchService{repo: repo}
}

func (s *MatchService) CreateMatch(match *domain.Match) error {
	return s.repo.Create(match)
}

func (s *MatchService) GetMatches() ([]*domain.Match, error) {
	return s.repo.GetAll()
}

func (s *MatchService) GetMatchByID(id int) (*domain.Match, error) {
	return s.repo.GetByID(id)
}

func (s *MatchService) UpdateMatch(match *domain.Match) error {
	return s.repo.Update(match)
}
