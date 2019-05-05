package betmodels

type Bet struct {
	BetID	string
	MatchID	string
	SelectedHomeTeam bool
	IsResolved bool
	Won	bool
}

type UserBets struct {
	UserID string
	Bets []*Bet
}
