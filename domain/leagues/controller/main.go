package leaguescontroller

import (
	"log"
	"net/http"
	// "encoding/json"
	// "github.com/gorilla/mux"
	// LeagueModels "github.com/alindenberg/know-it-all/domain/leagues/models"
)
var COLLECTION = "leagues"

func GetLeagues(w http.ResponseWriter, req *http.Request) {
	log.Println("Get Leagues")
}

func GetLeague(w http.ResponseWriter, req *http.Request) {
	log.Println("Get League")
}

func CreateLeague(w http.ResponseWriter, req *http.Request) {
	log.Println("Create League")
	// var league LeagueModels.League
	// decoder := json.NewDecoder(req.Body)
	// err := decoder.Decode(&league)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// league.LeagueID = uuid.New().String()
	// Repository.CreateMatch(league)
	
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// response := LeagueModels.CreateResponse{league.LeagueID}
	// json.NewEncoder(w).Encode(response)
}