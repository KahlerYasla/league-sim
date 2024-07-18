package domain

type Match struct {
	ID            int
	HomeTeamID    int
	AwayTeamID    int
	HomeTeamScore int
	AwayTeamScore int
	Week          int
	IsPlayed      bool
}
