package matchservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	MatchRepository "github.com/alindenberg/know-it-all/domain/matches/repository"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
	"github.com/google/uuid"
)

func CreateMatch(jsonBody io.ReadCloser) (string, error) {
	var match MatchModels.Match
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&match)
	if err != nil {
		return "", err
	}

	match.MatchID = uuid.New().String()
	match.HomeTeamScore = 0
	match.AwayTeamScore = 0

	err = validateMatch(&match)
	if err != nil {
		return "", err
	}
	return match.MatchID, MatchRepository.CreateMatch(&match)
}

func GetMatch(matchID string) (*MatchModels.Match, error) {
	_, err := uuid.Parse(matchID)
	if err != nil {
		return nil, err
	}

	match, err := MatchRepository.GetMatch(matchID)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func GetAllMatches(leagueID string, excludePast string) ([]*MatchModels.Match, error) {
	return MatchRepository.GetAllMatches(leagueID, excludePast)
}

func ResolveMatch(matchID string, jsonBody io.ReadCloser) error {
	var matchResult MatchModels.MatchResult
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&matchResult)
	if err != nil {
		return err
	}

	err = validateMatchResult(&matchResult)
	if err != nil {
		return err
	}

	err = MatchRepository.ResolveMatch(matchID, &matchResult)
	if err != nil {
		return err
	}

	err = UserService.ResolveBets(matchID, &matchResult)
	return err
}

func validateMatch(match *MatchModels.Match) error {
	_, err := uuid.Parse(match.LeagueID)
	if err != nil {
		return errors.New(fmt.Sprintf("Error decoding league id : %s", match.LeagueID))
	}
	if match.Date.Before(time.Now().UTC()) {
		return errors.New(fmt.Sprintf("New match must have Date after the current time"))
	}
	return nil
}

func validateMatchResult(matchResult *MatchModels.MatchResult) error {
	if matchResult.AwayScore < 0 {
		return errors.New("Away team score must be equal to or greater than 0")
	}
	if matchResult.HomeScore < 0 {
		return errors.New("Home team score must be equal to or greater than 0")
	}

	return nil
}
