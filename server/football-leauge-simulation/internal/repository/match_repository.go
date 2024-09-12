package repository

import (
	"database/sql"
	"football-league-simulation/internal/domain"
)

var matchColumns = []string{"id", "home_team_id", "away_team_id", "home_team_score", "away_team_score", "week", "is_played"}

func NewMatchRepository(db *sql.DB) Repository[domain.Match] {
	return NewGenericRepository[domain.Match](db, "matches", matchColumns)
}
