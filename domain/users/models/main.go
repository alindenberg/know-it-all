package usermodels

type UserCredentials struct {
	Username string
	Password string
	Email    string
}

type UserSignInRequest struct {
	Username string
	Password string
}

type UserKeys struct {
	Username    string
	AccessToken string
	RenewToken  string
}
type User struct {
	UserID string
	Bets   []UserBet
}

type UserBetRequest struct {
	MatchID    string
	Prediction Prediction
}

type UserBet struct {
	MatchID    string
	Prediction Prediction
	IsResolved bool
	Won        bool
}

type Prediction int

const (
	HomeTeam Prediction = 0
	AwayTeam Prediction = 1
	Draw     Prediction = 2
)
