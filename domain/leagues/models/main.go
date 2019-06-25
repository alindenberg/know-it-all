package leaguemodels

import "time"

type League struct {
	LeagueID string
	Name     string
	Country  string
	Division int
	Matches  []LeagueMatch
}

type AddMatchRequest struct {
	HomeTeam string
	AwayTeam string
	Date     string
}

type LeagueMatch struct {
	MatchID   string
	HomeTeam  string
	AwayTeam  string
	HomeScore int
	AwayScore int
	Date      time.Time
}

type LeagueMatchResult struct {
	HomeScore int
	AwayScore int
}
