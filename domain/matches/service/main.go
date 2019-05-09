package matchservice

import (
	"io"
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
	var match MatchModels.Match
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&match)
	if err != nil {
		return "", err
	}

	// TODO: Validation

	id := uuid.New().String()
	match.MatchID = id

	return id, MatchRepository.CreateMatch(match)
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
