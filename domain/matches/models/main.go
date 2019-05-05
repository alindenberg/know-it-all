package matchmodels

type Match struct {
	MatchID  string
	HomeTeam string
	AwayTeam string
	HomeScore	int
	AwayScore int
	Date string
}

type MatchResult struct {
	HomeScore	int
	AwayScore	int
}
