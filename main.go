package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	mongo "github.com/alindenberg/know-it-all/database"
	matchesController "github.com/alindenberg/know-it-all/domain/matches/controller"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
	groupsController "github.com/alindenberg/know-it-all/domain/groups/controller"
)

func main() {
	mongo.InitDatabase()
	addRouteHandlers()
	log.Println("Started Go server")
	startServer()
}

func addRouteHandlers() {
	r := httprouter.New()
	// MatchRoutes
	r.GET("/matches/:id", matchesController.GetMatch)
	r.GET("/matches", matchesController.GetAllMatches)
	r.POST("/matches", matchesController.CreateMatch)
	r.DELETE("/matches/:id", matchesController.DeleteMatch)

	// League Routes
	r.GET("/leagues/:id", leaguesController.GetLeague)
	r.GET("/leagues", leaguesController.GetAllLeagues)
	r.POST("/leagues", leaguesController.CreateLeague)
	r.DELETE("/leagues/:id", leaguesController.DeleteLeague)

	// Group Routes
	r.GET("/groups/:id", groupsController.GetGroup)
	r.GET("/groups", groupsController.GetAllGroups)
	r.POST("/groups", groupsController.CreateGroup)
	r.DELETE("/groups/:id", groupsController.DeleteGroup)

	// Register routes
	http.Handle("/", r)
}

func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}