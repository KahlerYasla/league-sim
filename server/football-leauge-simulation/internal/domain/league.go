package domain

type League struct {
	ID      int
	Name    string
	Teams   []*Team
	Matches []*Match
}
