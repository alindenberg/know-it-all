package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mongo "github.com/alindenberg/know-it-all/database"
	matchesController "github.com/alindenberg/know-it-all/domain/matches/controller"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
)

func main() {
	mongo.InitDatabase()
	addRouteHandlers()
	log.Println("Started Go server")
	startServer()
}

func addRouteHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/matches", matchesHandler)
	r.HandleFunc("/matches/{id}", matchHandler)
	r.HandleFunc("/leagues", leaguesHandler)
	r.HandleFunc("/leagues/{id}", leagueHandler)
	http.Handle("/", r)
}

func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func matchesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		matchesController.GetAllMatches(w, req)
		break
	case http.MethodPost:
		matchesController.CreateMatch(w, req)
		break
	default:
		log.Println(w, "Application can't handle "+req.Method+" requests")
	}
}
func matchHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		matchesController.GetMatch(w, req)
		break
	default:
		break
	}
}

func leaguesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		leaguesController.GetLeagues(w, req)
		break
	case http.MethodPost:
		leaguesController.CreateLeague(w, req)
		break
	default:
		log.Println(w, "Application can't handle "+req.Method+" requests")
	}
}
func leagueHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		leaguesController.GetLeague(w, req)
		break
	default:
		break
	}
}
