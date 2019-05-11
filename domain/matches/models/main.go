package matchmodels

import (
	"time"
)

type MatchRequest struct {
	HomeTeam string
	AwayTeam string
	Date string
}

type Match struct {
	MatchID  string
	HomeTeam string
	AwayTeam string
	HomeScore	int
	AwayScore int
	Date time.Time
}

type MatchResult struct {
	HomeScore	int
	AwayScore	int
}
