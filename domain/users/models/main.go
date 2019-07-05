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
type UserRequest struct {
	UserID string
	Email  string
}
type User struct {
	UserID        string
	Email         string
	Username      string
	Bets          []UserBet
	Friends       []string
	Wins          int
	Losses        int
	WinPercentage float64
}

type CreateBetRequest struct {
	MatchID    string
	LeagueID   string
	Prediction Prediction
}

type CreateUsernameRequest struct {
	Username string
}

type UpdateBetRequest struct {
	Prediction Prediction
}

type UserBet struct {
	MatchID    string
	LeagueID   string
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
