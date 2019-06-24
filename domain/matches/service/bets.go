package matchservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	Models "github.com/alindenberg/know-it-all/domain/matches/models"
	Repository "github.com/alindenberg/know-it-all/domain/matches/repository"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
	"github.com/google/uuid"
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

func GetAllBets(queryMap map[string][]string) ([]*Models.Bet, error) {
	var results []*Models.Bet
	var err error

	if queryMap["userId"] != nil {
		userId := queryMap["userId"][0]
		_, err = uuid.Parse(userId)
		if err != nil {
			return nil, err
		}
		results, err = Repository.GetAllBetsForUser(queryMap["userId"][0])
	} else if queryMap["matchId"] != nil {
		matchId := queryMap["matchId"][0]
		_, err = uuid.Parse(matchId)
		if err != nil {
			return nil, err
		}
		results, err = Repository.GetAllBetsForMatch(matchId)
	} else {
		results, err = Repository.GetAllBets()
	}

	return results, err
}

func GetAllBetsForMatch(matchId string) ([]*Models.Bet, error) {
	return Repository.GetAllBetsForMatch(matchId)
}

func GetAllBetsForUser(userId string) ([]*Models.Bet, error) {
	return Repository.GetAllBetsForUser(userId)
}

func CreateBet(jsonBody io.ReadCloser) (string, error) {
	var betRequest Models.BetRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&betRequest)
	if err != nil {
		return "", err
	}

	bet := betFromRequest(&betRequest)

	// Validate group properties
	err = validateBet(bet)
	if err != nil {
		return "", err
	}

	return bet.BetID, Repository.CreateBet(bet)
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
			if result.HomeScore > result.AwayScore {
				wonBet = true
			} else {
				wonBet = false
			}
			Repository.ResolveBet(bet.BetID, wonBet)
			break
		case Models.AwayTeam:
			if result.HomeScore < result.AwayScore {
				wonBet = true
			} else {
				wonBet = false
			}
			Repository.ResolveBet(bet.BetID, wonBet)
			break
		case Models.Draw:
			if result.HomeScore == result.AwayScore {
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
		return errors.New(fmt.Sprintf("MatchId : ", err))
	}

	_, err = uuid.Parse(bet.UserID)
	if err != nil {
		return errors.New(fmt.Sprintf("UserId : ", err))
	}

	// validate userId corresponds to existing user
	_, err = UserService.GetUser(bet.UserID)
	if err != nil {
		return errors.New(fmt.Sprintf("No corresponding User found with id: %s", bet.UserID))
	}

	// validate matchId corresponds to existing match
	match, err := GetMatch(bet.MatchID)
	if err != nil {
		return errors.New(fmt.Sprintf("No corresponding Match found with id: %s", bet.MatchID))
	}

	// Validate match has not taken place yet
	if match.Date.Before(time.Now().UTC()) {
		return errors.New(fmt.Sprintf("May not place a bet on a Match that has begun or completed."))
	}

	// Validate valid result selection (TeamSelection enum)
	if bet.Prediction != Models.HomeTeam && bet.Prediction != Models.AwayTeam && bet.Prediction != Models.Draw {
		return errors.New(
			fmt.Sprintf("Invalid Prediction. Must be choice of Home Team (0), Away Team (1), or Draw (2)"))
	}

	return nil
}

func betFromRequest(betRequest *Models.BetRequest) *Models.Bet {
	return &Models.Bet{
		uuid.New().String(),
		betRequest.MatchID,
		betRequest.UserID,
		betRequest.Prediction,
		false,
		false,
	}
}
