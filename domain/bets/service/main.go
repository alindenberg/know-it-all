package betservice

import (
	"io"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	BetModels "github.com/alindenberg/know-it-all/domain/bets/models"
	BetRepository "github.com/alindenberg/know-it-all/domain/bets/repository"
	// MatchService "github.com/alindenberg/know-it-all/domain/matches/service"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
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

func GetAllBetsForUser(userId string) ([]*BetModels.Bet, error) {
	return BetRepository.GetAllBetsForUser(userId)
}

func GetAllBetsForMatch(matchId string) ([]*BetModels.Bet, error) {
	return BetRepository.GetAllBetsForMatch(matchId)
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

func ResolveBets(matchId string, result *MatchModels.MatchResult) error {
	bets, err := BetRepository.GetAllBetsForMatch(matchId)
	if err != nil {
		return err
	}

	for _, bet := range bets {
		switch bet.Selection {
		case BetModels.HomeTeam:
			if(result.HomeScore > result.AwayScore) {
				log.Println("About to repo resolve bet")
				return BetRepository.ResolveBet(bet.BetID, true)
			} else {
				return BetRepository.ResolveBet(bet.BetID, false)
			}
			break
		case BetModels.AwayTeam:
			if(result.HomeScore < result.AwayScore) {
				log.Println("About to repo resolve bet")
				return BetRepository.ResolveBet(bet.BetID, true)
			} else {
				return BetRepository.ResolveBet(bet.BetID, false)
			}
			break
		case BetModels.Draw:
			if(result.HomeScore == result.AwayScore) {
				log.Println("About to repo resolve bet")
				return BetRepository.ResolveBet(bet.BetID, true)
			} else {
				return BetRepository.ResolveBet(bet.BetID, false)
			}
			break
		default:
			log.Println("No Selection made")
			break
		}
	}

	return nil
}

func validateBet(bet *BetModels.Bet) error {
	// validate matchId corresponds to existing match
	// _, err := MatchService.GetMatch(bet.MatchID)
	// if err != nil {
	// 	return errors.New(fmt.Sprintf("No match found with id: %s", bet.MatchID))
	// }

	return nil
}
