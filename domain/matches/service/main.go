package matchservice

import (
	"io"
	"fmt"
	"log"
	"time"
	"errors"
	"encoding/json"
	"github.com/google/uuid"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	MatchRepository "github.com/alindenberg/know-it-all/domain/matches/repository"
)

func GetMatch(id string) (*MatchModels.Match, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return MatchRepository.GetMatch(id)
}

func GetAllMatches() ([]*MatchModels.Match, error) {
	return  MatchRepository.GetAllMatches()
}

func CreateMatch(jsonBody io.ReadCloser) (string, error) {
	var matchRequest MatchModels.MatchRequest
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&matchRequest)
	if err != nil {
		return "", err
	}

	match, err := matchFromRequest(&matchRequest)
	if err != nil {
		return "", err
	}
	err = validateMatch(match)
	if err != nil {
		return "", err
	}

	return match.MatchID, MatchRepository.CreateMatch(match)
}

func DeleteMatch(id string) error {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return MatchRepository.DeleteMatch(id)
}

func ResolveMatch(id string, jsonBody io.ReadCloser) error {
	var matchResult MatchModels.MatchResult

	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&matchResult)
	if err != nil {
		return err
	}

	err = MatchRepository.ResolveMatch(id, &matchResult)
	if err != nil {
		return err
	}
	return ResolveBets(id, &matchResult)
}

func matchFromRequest(matchRequest *MatchModels.MatchRequest) (*MatchModels.Match, error) {
	// Create datetime field from string
	datetime, err := time.Parse(time.RFC3339, matchRequest.Date)
	if err != nil {
		return nil, err
	}

	return &MatchModels.Match{uuid.New().String(), matchRequest.HomeTeam, matchRequest.AwayTeam, 0, 0, datetime}, nil
}

func validateMatch(match *MatchModels.Match) error {
	// Max Length checks for team names
	if len(match.HomeTeam) > 20 {
		return errors.New(fmt.Sprintf("Home Team name may not be longer than 20 characters"))
	}
	if len(match.HomeTeam) > 20 {
		return errors.New(fmt.Sprintf("Away Team name may not be longer than 20 characters"))
	}
	// Validate scores are init to 0
	if match.HomeScore != 0 {
		return errors.New(fmt.Sprintf("Home Score must be 0 for match initilization"))
	}
	if match.AwayScore != 0 {
		return errors.New(fmt.Sprintf("Away Score must be 0 for match initilization"))
	}
	log.Println(time.Now())
	// validate that match datetime is in the future
	if match.Date.Before(time.Now().UTC()) {
		return errors.New(fmt.Sprintf("New match datetime may not be in the past"))
	}

	return nil
}
