package leaderboardservice

import (
	LeaderboardModels "github.com/alindenberg/know-it-all/domain/leaderboard/models"
)

func GetLeaderboard() ([]*LeaderboardModels.LeaderboardEntry, error) {
	return LeaderboardRepository.GetLeaderboard()
}

func GetLeaderboardForUser(userAndFriendIds []string) ([]*LeaderboardModels.LeaderboardEntry, error) {
	return LeaderboardRepository.GetLeaderboardOnlyForIds()
}
