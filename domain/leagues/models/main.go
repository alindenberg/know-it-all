package leaguemodels

import (
	"time"
)

type League struct {
	LeagueID string
	Name     string
	LogoURL  string
	Country  string
	Division int
	// Embedded array of league's matches
	Matches []LeagueMatch
}

type LeagueMatch struct {
	MatchID       string
	HomeTeamID    string
	AwayTeamID    string
	HomeTeamScore int
	AwayTeamScore int
	Date          time.Time
}

type LeagueMatchResult struct {
	HomeScore int
	AwayScore int
}
