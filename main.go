package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	mongo "github.com/alindenberg/know-it-all/database"
	matchesController "github.com/alindenberg/know-it-all/domain/matches/controller"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
	usersController "github.com/alindenberg/know-it-all/domain/users/controller"
	betsController "github.com/alindenberg/know-it-all/domain/bets/controller"
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
	r.POST("/matches/:id/resolve", matchesController.ResolveMatch)
	r.DELETE("/matches/:id", matchesController.DeleteMatch)

	// League Routes
	r.GET("/leagues/:id", leaguesController.GetLeague)
	r.GET("/leagues", leaguesController.GetAllLeagues)
	r.POST("/leagues", leaguesController.CreateLeague)
	r.DELETE("/leagues/:id", leaguesController.DeleteLeague)

	// User Routes
	r.GET("/users/:id", usersController.GetUser)
	r.GET("/users", usersController.GetAllUsers)
	r.POST("/users", usersController.CreateUser)
	r.DELETE("/users/:id", usersController.DeleteUser)

	// Bet Routes
	// r.GET("/users/:id/bets/:betId", betsController.GetBet)
	r.GET("/users/:id/bets", betsController.GetAllBets)
	r.POST("/users/:id/bets", betsController.CreateBet)
	r.DELETE("/users/:id/bets/:betId", betsController.DeleteBet)


	// Register routes
	http.Handle("/", r)
}

func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
