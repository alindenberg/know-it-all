package matchescontroller

import (
	"encoding/json"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	Repository "github.com/alindenberg/know-it-all/domain/matches/repository"
)

var COLLECTION = "matches"

func GetMatch(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	result, err := Repository.GetMatch(id)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMatches(w http.ResponseWriter, req *http.Request) {
	results, err := Repository.GetAllMatches()
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateMatch(w http.ResponseWriter, req *http.Request) {
	var match MatchModels.Match
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&match)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	match.MatchID = uuid.New().String()
	
	Repository.CreateMatch(match)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, match.MatchID)
}
