package main

import (
	"log"
	"net/http"

	mongo "github.com/alindenberg/know-it-all/database"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
	matchesController "github.com/alindenberg/know-it-all/domain/matches/controller"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	usersController "github.com/alindenberg/know-it-all/domain/users/controller"
	userService "github.com/alindenberg/know-it-all/domain/users/service"
	"github.com/julienschmidt/httprouter"
	// betsController "github.com/alindenberg/know-it-all/domain/bets/controller"
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
	r.POST("/users/sessions/create", usersController.CreateUserSession)
	r.DELETE("/users/:id", usersController.DeleteUser)

	// Bet routes
	r.GET("/bets", BasicAuth(matchesController.GetAllBets))
	r.POST("/bets", matchesController.CreateBet)
	r.DELETE("/bets/:betId", matchesController.DeleteBet)

	// Register routes
	http.Handle("/", r)
}

func BasicAuth(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		err := userService.Authenticate(r)
		if err != nil {
			SharedResponses.Error(w, err)
			return
		}
		handler(w, r, ps)
	}

}

func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
