package matchmodels

type BetRequest struct {
	MatchID string
	Prediction TeamSelection
}
type Bet struct {
	BetID	string
	MatchID	string
	UserID string
	Prediction TeamSelection
	IsResolved bool
	Won bool
}

type TeamSelection int

const (
	HomeTeam TeamSelection = 0
	AwayTeam TeamSelection = 1
	Draw TeamSelection = 2
)
