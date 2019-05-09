package matchservice

import (
	"io"
	"log"
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

	log.Println(bets)
	for _, bet := range bets {
		wonBet := false

		switch bet.Selection {
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
	// validate matchId corresponds to existing match
	// _, err := MatchService.GetMatch(bet.MatchID)
	// if err != nil {
	// 	return errors.New(fmt.Sprintf("No match found with id: %s", bet.MatchID))
	// }

	return nil
}
