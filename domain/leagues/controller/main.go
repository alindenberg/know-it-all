package leaguescontroller

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	LeagueService "github.com/alindenberg/know-it-all/domain/leagues/service"
)

func GetLeague(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

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

func DeleteLeague(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := LeagueService.DeleteLeague(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Delete(w)
}
