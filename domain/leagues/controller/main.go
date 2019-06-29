package leaguescontroller

import (
	"encoding/json"
	"net/http"

	LeagueService "github.com/alindenberg/know-it-all/domain/leagues/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	"github.com/julienschmidt/httprouter"
)

func GetLeague(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("leagueId")

	result, err := LeagueService.GetLeague(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllLeagues(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	results, err := LeagueService.GetAllLeagues()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateLeague(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	id, err := LeagueService.CreateLeague(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func CreateLeagueMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	leageuID := params.ByName("leagueId")

	id, err := LeagueService.CreateLeagueMatch(leageuID, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}

func ResolveLeagueMatch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	leagueID := params.ByName("leagueId")
	matchID := params.ByName("matchId")

	err := LeagueService.ResolveLeagueMatch(leagueID, matchID, req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)

}

func DeleteLeague(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := LeagueService.DeleteLeague(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.NoContent(w)
}
