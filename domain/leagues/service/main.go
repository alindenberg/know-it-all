package leagueservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	LeagueRepository "github.com/alindenberg/know-it-all/domain/leagues/repository"
	UserService "github.com/alindenberg/know-it-all/domain/users/service"
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
	// initialize matches to empty array
	league.UpcomingMatches = []LeagueModels.LeagueMatch{}
	league.PastMatches = []LeagueModels.LeagueMatch{}

	err = validateLeague(&league)
	if err != nil {
		return "", err
	}

	return league.LeagueID, LeagueRepository.CreateLeague(league)
}

func GetLeagueMatch(leagueID string, matchID string) (*LeagueModels.LeagueMatch, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(leagueID)
	if err != nil {
		return nil, err
	}

	_, err = uuid.Parse(matchID)
	if err != nil {
		return nil, err
	}

	matches, err := LeagueRepository.GetLeagueMatches(leagueID)
	if err != nil {
		return nil, err
	}

	for _, match := range *matches {
		if match.MatchID == matchID {
			return &match, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("No match found in leagueId %s for matchId %s", leagueID, matchID))
}

func GetAllLeagueMatches(leagueID string) (*[]LeagueModels.LeagueMatch, error) {
	return LeagueRepository.GetLeagueMatches(leagueID)
}

func CreateLeagueMatch(leagueId string, jsonBody io.ReadCloser) (string, error) {
	var match LeagueModels.LeagueMatch
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
	return match.MatchID, LeagueRepository.CreateLeagueMatch(leagueId, &match)
}

func ResolveLeagueMatch(leagueID string, matchID string, jsonBody io.ReadCloser) error {
	var matchResult LeagueModels.LeagueMatchResult
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&matchResult)
	if err != nil {
		return err
	}

	err = validateMatchResult(&matchResult)
	if err != nil {
		return err
	}

	err = LeagueRepository.ResolveLeagueMatch(leagueID, matchID, &matchResult)
	if err != nil {
		return err
	}

	err = UserService.ResolveBets(matchID, &matchResult)
	return err
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
		return errors.New(fmt.Sprintf("Field 'name' may not be longer than 25 characters"))
	}

	if len(league.Country) > 25 {
		return errors.New(fmt.Sprintf("Field 'country' name may not be longer than 25 characters. Use abbreviation if necessary."))
	}

	if league.Division <= 0 {
		return errors.New(fmt.Sprintf("Field 'division' must be an integer greater than 0"))
	}

	return nil
}

func validateMatch(match *LeagueModels.LeagueMatch) error {
	if match.Date.Before(time.Now().UTC()) {
		return errors.New(fmt.Sprintf("New match must have Date after the current time"))
	}
	return nil
}

func validateMatchResult(matchResult *LeagueModels.LeagueMatchResult) error {
	if matchResult.AwayScore < 0 {
		return errors.New("Away team score must be equal to or greater than 0")
	}
	if matchResult.HomeScore < 0 {
		return errors.New("Home team score must be equal to or greater than 0")
	}

	return nil
}
