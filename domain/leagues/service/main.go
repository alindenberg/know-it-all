package leagueservice

import (
	"io"
	"fmt"
	"errors"
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

	id := uuid.New().String()
	league.LeagueID = id

	err = validateLeague(&league)
	if err != nil {
		return "", err
	}

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

func validateLeague(league *LeagueModels.League) error {
	_, err := uuid.Parse(league.LeagueID)
	if err != nil {
		return errors.New(fmt.Sprintf("LeagueId : ", err))
	}

	if len(league.Name) > 25 {
		return errors.New(fmt.Sprintf("League name may not be longer than 25 characters"))
	}

	if len(league.Country) > 25 {
		return errors.New(fmt.Sprintf("League country name may not be longer than 25 characters. Use abbreviation if necessary."))
	}

	if league.Division <= 0 {
		return errors.New(fmt.Sprintf("Division must be an integer greater than 0"))
	}

	return nil
}
