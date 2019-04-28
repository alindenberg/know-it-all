package matchescontroller

import (
	"log"

	// "bytes"

	// "io/ioutil"
	"encoding/json"
	"net/http"

	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	Repository "github.com/alindenberg/know-it-all/domain/matches/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var COLLECTION = "matches"

func GetMatch(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	// Minimal input sanitization on id value
	// just make sure its valid uuid
	_, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
	}

	result := Repository.GetMatch(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func GetAllMatches(w http.ResponseWriter, req *http.Request) {
	var results = Repository.GetAllMatches()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

func CreateMatch(w http.ResponseWriter, req *http.Request) {
	var match MatchModels.Match
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&match)
	if err != nil {
		log.Fatal(err)
	}

	match.MatchID = uuid.New().String()
	Repository.CreateMatch(match)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := MatchModels.CreateResponse{match.MatchID}
	json.NewEncoder(w).Encode(response)
}
