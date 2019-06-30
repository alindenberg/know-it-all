package leaderboardmodels

type LeaderboardEntry struct {
	UserID        string
	Wins          int64
	Losses        int64
	WinPercentage float64
}
