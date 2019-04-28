package matchmodels

type Match struct {
	MatchID  string
	HomeTeam string
	AwayTeam string
	Date     string
	// Home = 0, Away = 1
	Winner string
}