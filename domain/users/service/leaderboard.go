package userservice

import (
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	UserRepository "github.com/alindenberg/know-it-all/domain/users/repository"
)

func GetLeaderboard() ([]*UserModels.User, error) {
	return UserRepository.GetLeaderboard()
}

func GetLeaderboardForUser(userID string) ([]*UserModels.User, error) {
	user, err := UserRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	usersNeeded := append(user.Friends, user.UserID)
	return UserRepository.GetLeaderboardOnlyForIds(usersNeeded)
}
