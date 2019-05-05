package betservice

import (
	"io"
	"encoding/json"
	"github.com/google/uuid"
	BetModels "github.com/alindenberg/know-it-all/domain/bets/models"
	BetRepository "github.com/alindenberg/know-it-all/domain/bets/repository"
)

// func GetBet(id string, userId string) (*BetModels.Bet, error) {
// 	// Minimal input sanitization on id value
// 	// just make sure its valid uuid
// 	_, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return BetRepository.GetBet(id)
// }

func GetAllBets(userId string) ([]*BetModels.Bet, error) {
	return BetRepository.GetAllBets(userId)
}

func CreateBet(jsonBody io.ReadCloser, userId string) (string, error) {
	var bet BetModels.Bet
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&bet)
	if err != nil {
		return "", err
	}

	// Assign a random id
	id := uuid.New().String()
	bet.BetID = id

	// Validate group properties
	err = validateBet(&bet)

	if err != nil {
		return "", err
	}

	return id, BetRepository.CreateBet(bet, userId)
}

func DeleteBet(id string, userId string) error {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = uuid.Parse(userId)
	if err != nil {
		return err
	}

	return BetRepository.DeleteBet(id, userId)
}

func validateBet(bet *BetModels.Bet) error {
	return nil
}
