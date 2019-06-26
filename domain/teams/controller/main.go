package teamcontroller

import (
	"encoding/json"
	"net/http"

	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	TeamService "github.com/alindenberg/know-it-all/domain/teams/service"
	"github.com/julienschmidt/httprouter"
)

// GetTeam - get team by id
func GetTeam(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	teamID := params.ByName("teamId")

	result, err := TeamService.GetTeam(teamID)

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetAllTeams - get all teams
func GetAllTeams(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	result, err := TeamService.GetAllTeams()

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllTeamsForLeague(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	leagueID := params.ByName("leagueId")
	result, err := TeamService.GetAllTeamsForLeague(leagueID)

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateTeam(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	teamID, err := TeamService.CreateTeam(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, teamID)
}

// func DeleteMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
// 	id := params.ByName("id")

// 	err := MatchService.DeleteMatch(id)
// 	if err != nil {
// 		SharedResponses.Error(w, err)
// 		return
// 	}

// 	SharedResponses.NoContent(w)
// }
