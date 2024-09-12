package repository

import (
	"database/sql"
	"football-league-simulation/internal/domain"
)

var leagueColumns = []string{"id", "name"}

func NewLeagueRepository(db *sql.DB) Repository[domain.League] {
	return NewGenericRepository[domain.League](db, "leagues", leagueColumns)
}
