package domain

type Team struct {
	ID             int
	Name           string
	Strength       int
	Played         int
	Won            int
	Drawn          int
	Lost           int
	Points         int
	GoalDifference int
}
