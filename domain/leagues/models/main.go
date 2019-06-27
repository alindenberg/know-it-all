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
	// Embedded array of league's upcoming matches
	UpcomingMatches []LeagueMatch
	// Embedded array of league's past matches
	PastMatches []LeagueMatch
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
