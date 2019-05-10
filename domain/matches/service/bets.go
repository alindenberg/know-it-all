package matchservice

import (
	"io"
	"errors"
	"fmt"
	"encoding/json"
	"github.com/google/uuid"
	Models "github.com/alindenberg/know-it-all/domain/matches/models"
	Repository "github.com/alindenberg/know-it-all/domain/matches/repository"
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

func GetAllBetsForUser(userId string) ([]*Models.Bet, error) {
	return Repository.GetAllBetsForUser(userId)
}

func GetAllBetsForMatch(matchId string) ([]*Models.Bet, error) {
	return Repository.GetAllBetsForMatch(matchId)
}

func CreateBet(jsonBody io.ReadCloser, userId string) (string, error) {
	var bet Models.Bet
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&bet)
	if err != nil {
		return "", err
	}

	// Assign a random id
	id := uuid.New().String()
	bet.BetID = id

	// Assign userId from request path
	bet.UserID = userId

	// Validate group properties
	err = validateBet(&bet)
	if err != nil {
		return "", err
	}

	return id, Repository.CreateBet(bet, userId)
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

	return Repository.DeleteBet(id, userId)
}

func ResolveBets(matchId string, result *Models.MatchResult) error {
	bets, err := Repository.GetAllBetsForMatch(matchId)
	if err != nil {
		return err
	}

	for _, bet := range bets {
		wonBet := false

		switch bet.Prediction {
		case Models.HomeTeam:
			if(result.HomeScore > result.AwayScore) {
				wonBet = true
			} else {
				wonBet = false
			}
			Repository.ResolveBet(bet.BetID, wonBet)
			break
		case Models.AwayTeam:
			if(result.HomeScore < result.AwayScore) {
				wonBet = true
			} else {
				wonBet = false
			}
			Repository.ResolveBet(bet.BetID, wonBet)
			break
		case Models.Draw:
			if(result.HomeScore == result.AwayScore) {
				wonBet = true
			} else {
				wonBet = false
			}
			Repository.ResolveBet(bet.BetID, wonBet)
			break
		default:
			break
		}
	}

	return nil
}

func validateBet(bet *Models.Bet) error {
	_, err := uuid.Parse(bet.MatchID)
	if err != nil {
		return err
	}

	_, err = uuid.Parse(bet.UserID)
	if err != nil {
		return err
	}

	// validate matchId corresponds to existing match
	_, err = GetMatch(bet.MatchID)
	if err != nil {
		return errors.New(fmt.Sprintf("No corresponding Match found with id: %s", bet.MatchID))
	}

	// TODO : Once datetime functionality is on matches
	// Validate the match isn't already resolved
	// if match.Datetime >= time.Now()  {
	// 	return errors.New(fmt.Sprintf("Can't place a bet on a match that's already begun or completed."))
	// }

	// Validate valid result selection (TeamSelection enum)
	if bet.Prediction != Models.HomeTeam && bet.Prediction != Models.AwayTeam && bet.Prediction != Models.Draw {
		return errors.New(
			fmt.Sprintf("Invalid Prediction. Must be choice of Home Team (0), Away Team (1), or Draw (2)"))
	}

	// Reset logic fields to be false if they were sent as true
	if bet.IsResolved {
		bet.IsResolved = false
	}
	if bet.Won {
		bet.Won = false
	}

	return nil
}
