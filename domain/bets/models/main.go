package betmodels

type Bet struct {
	BetID	string
	MatchID	string
	UserID string
	Selection TeamSelection
	IsResolved bool
	Won bool
}

type TeamSelection int

const (
	HomeTeam TeamSelection = 0
	AwayTeam TeamSelection = 1
	Draw TeamSelection = 2
)
