package matchmodels

import "time"

type Match struct {
	MatchID       string
	LeagueID      string
	HomeTeamID    string
	AwayTeamID    string
	HomeTeamScore int
	AwayTeamScore int
	Date          time.Time
	IsResolved    bool
}

type MatchResult struct {
	HomeScore int
	AwayScore int
}
