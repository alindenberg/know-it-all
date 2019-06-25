package leagueservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	LeagueRepository "github.com/alindenberg/know-it-all/domain/leagues/repository"
	"github.com/google/uuid"
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

	league.LeagueID = uuid.New().String()

	err = validateLeague(&league)
	if err != nil {
		return "", err
	}

	return league.LeagueID, LeagueRepository.CreateLeague(league)
}

func CreateLeagueMatch(leagueId string, jsonBody io.ReadCloser) (string, error) {
	var match LeagueModels.LeagueMatch
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&match)
	if err != nil {
		return "", err
	}

	match.MatchID = uuid.New().String()

	err = validateMatch(&match)
	if err != nil {
		return "", err
	}
	return match.MatchID, LeagueRepository.CreateLeagueMatch(leagueId, &match)
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

func validateMatch(match *LeagueModels.LeagueMatch) error {
	if match.Date.Before(time.Now().UTC()) {
		return errors.New(fmt.Sprintf("New match must have Date after the current time"))
	}
	return nil
}
