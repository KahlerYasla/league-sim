package repository

import (
	"database/sql"
	"football-league-simulation/internal/domain"
)

var teamColumns = []string{"id", "name", "strength", "points", "played", "won", "drawn", "lost", "goal_diff"}

func NewTeamRepository(db *sql.DB) Repository[domain.Team] {
	return NewGenericRepository[domain.Team](db, "teams", teamColumns)
}
