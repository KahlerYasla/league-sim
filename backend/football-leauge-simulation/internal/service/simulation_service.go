package service

import (
	"football-league-simulation/internal/domain"
	"football-league-simulation/internal/util"
	"log"
)

type SimulationService struct {
	teamService   *TeamService
	matchService  *MatchService
	leagueService *LeagueService
}

func NewSimulationService(teamService *TeamService, matchService *MatchService, leagueService *LeagueService) *SimulationService {
	return &SimulationService{
		teamService:   teamService,
		matchService:  matchService,
		leagueService: leagueService,
	}
}

func (s *SimulationService) SimulateWeek(week int) error {
	teams, err := s.teamService.GetTeams()
	if err != nil {
		return err
	}

	for _, homeTeam := range teams {
		for _, awayTeam := range teams {
			if homeTeam.ID != awayTeam.ID {
				match := &domain.Match{
					HomeTeamID: homeTeam.ID,
					AwayTeamID: awayTeam.ID,
					Week:       week,
					IsPlayed:   true,
				}
				homeTeamScore := util.RandomInt(0, homeTeam.Strength)
				awayTeamScore := util.RandomInt(0, awayTeam.Strength)
				match.HomeTeamScore = homeTeamScore
				match.AwayTeamScore = awayTeamScore

				if homeTeamScore > awayTeamScore {
					homeTeam.Won++
					homeTeam.Points += 3
					awayTeam.Lost++
				} else if homeTeamScore < awayTeamScore {
					awayTeam.Won++
					awayTeam.Points += 3
					homeTeam.Lost++
				} else {
					homeTeam.Drawn++
					awayTeam.Drawn++
					homeTeam.Points++
					awayTeam.Points++
				}

				homeTeam.Played++
				awayTeam.Played++
				homeTeam.GoalDifference += homeTeamScore - awayTeamScore
				awayTeam.GoalDifference += awayTeamScore - homeTeamScore

				if err := s.matchService.CreateMatch(match); err != nil {
					return err
				}

				if err := s.teamService.UpdateTeam(homeTeam); err != nil {
					return err
				}

				if err := s.teamService.UpdateTeam(awayTeam); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *SimulationService) SimulateSeason() error {
	for week := 1; week <= 4; week++ { // Assuming a 4-week season
		if err := s.SimulateWeek(week); err != nil {
			log.Printf("Error simulating week %d: %v", week, err)
			return err
		}
	}
	return nil
}
