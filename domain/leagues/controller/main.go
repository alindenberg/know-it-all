package leaguescontroller

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	LeagueService "github.com/alindenberg/know-it-all/domain/leagues/service"
)
var COLLECTION = "leagues"

func GetLeague(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	result, err := LeagueService.GetLeague(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllLeagues(w http.ResponseWriter, req *http.Request) {
	results, err := LeagueService.GetAllLeagues()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateLeague(w http.ResponseWriter, req *http.Request) {
	id, err := LeagueService.CreateLeague(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}