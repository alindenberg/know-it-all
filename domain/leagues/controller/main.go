package leaguescontroller

import (
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
	LeagueRepository "github.com/alindenberg/know-it-all/domain/leagues/repository"
)
var COLLECTION = "leagues"

func GetLeague(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	result, err := LeagueRepository.GetLeague(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllLeagues(w http.ResponseWriter, req *http.Request) {
	results, err := LeagueRepository.GetAllLeagues()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateLeague(w http.ResponseWriter, req *http.Request) {
	var league LeagueModels.League
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&league)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	league.LeagueID = uuid.New().String()
	err = LeagueRepository.CreateLeague(league)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, league.LeagueID)
}