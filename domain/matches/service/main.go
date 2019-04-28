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