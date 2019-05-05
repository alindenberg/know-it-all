package leagueservice

import (
	"io"
	"encoding/json"
	"github.com/google/uuid"
	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	LeagueRepository "github.com/alindenberg/know-it-all/domain/leagues/repository"
)

func GetLeague(id string) (*LeagueModels.League, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return LeagueRepository.GetLeague(id)
} 

func GetAllLeagues() ([]*LeagueModels.League, error) {
	return LeagueRepository.GetAllLeagues()
}

func CreateLeague(jsonBody io.ReadCloser) (string, error) {
	var league LeagueModels.League
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&league)
	if err != nil {
		return "", err
	}

	// TODO: Validation

	id := uuid.New().String()
	league.LeagueID = id

	return id, LeagueRepository.CreateLeague(league)
}

func DeleteLeague(id string) error {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return LeagueRepository.DeleteLeague(id)
}