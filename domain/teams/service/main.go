package teamservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	TeamModels "github.com/alindenberg/know-it-all/domain/teams/models"
	TeamRepository "github.com/alindenberg/know-it-all/domain/teams/repository"
	"github.com/google/uuid"
)

func GetTeam(teamID string) (*TeamModels.Team, error) {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(teamID)
	if err != nil {
		return nil, err
	}

	return TeamRepository.GetTeam(teamID)
}

func GetAllTeams() ([]*TeamModels.Team, error) {
	return TeamRepository.GetAllTeams()
}

func GetAllTeamsForLeague(leagueID string) ([]*TeamModels.Team, error) {
	return TeamRepository.GetAllTeamsForLeague(leagueID)
}

func CreateTeam(jsonBody io.ReadCloser) (string, error) {
	var team TeamModels.Team
	decoder := json.NewDecoder(jsonBody)
	err := decoder.Decode(&team)
	if err != nil {
		return "", err
	}

	team.TeamID = uuid.New().String()

	err = validateTeam(&team)
	if err != nil {
		return "", err
	}

	return team.TeamID, TeamRepository.CreateTeam(&team)
}

func DeleteTeam(teamID string) error {
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(teamID)
	if err != nil {
		return err
	}

	return TeamRepository.DeleteTeam(teamID)
}

func validateTeam(team *TeamModels.Team) error {
	// TODO: validate
	if len(team.Leagues) < 1 {
		return errors.New("Team must be associated with at least 1 league")
	}
	for _, leagueId := range team.Leagues {
		_, err := uuid.Parse(leagueId)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid league id : %s. %s", leagueId, err.Error()))
		}
	}
	return nil
}
