package matchescontroller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	MatchService "github.com/alindenberg/know-it-all/domain/matches/service"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
)

var COLLECTION = "matches"

func GetMatch(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	result, err := MatchService.GetMatch(id) 

	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMatches(w http.ResponseWriter, req *http.Request) {
	result, err := MatchService.GetAllMatches()
	
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateMatch(w http.ResponseWriter, req *http.Request) {
	
	id, err := MatchService.CreateMatch(req.Body)
	if err != nil {
		SharedResponses.Error(w, err)
		return
	}

	SharedResponses.Create(w, id)
}
