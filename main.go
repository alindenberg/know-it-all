package main

import (
	"errors"
	"log"
	"net/http"

	mongo "github.com/alindenberg/know-it-all/database"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
	SharedResponses "github.com/alindenberg/know-it-all/domain/shared/responses"
	teamsController "github.com/alindenberg/know-it-all/domain/teams/controller"
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
	// TeamRoutes
	r.GET("/teams", teamsController.GetAllTeams)
	r.GET("/teams/:teamId", teamsController.GetTeam)

	r.POST("/teams", teamsController.CreateTeam)

	// League Routes
	r.GET("/leagues", leaguesController.GetAllLeagues)
	r.GET("/leagues/:id", leaguesController.GetLeague)
	r.GET("/leagues/:id/teams", teamsController.GetAllTeamsForLeague)

	r.POST("/leagues", leaguesController.CreateLeague)
	r.POST("/leagues/:leagueId/matches", leaguesController.CreateLeagueMatch)

	r.DELETE("/leagues/:id", leaguesController.DeleteLeague)

	// User Routes
	r.GET("/users/:id", usersController.GetUser)
	r.GET("/users", usersController.GetAllUsers)

	r.POST("/users", usersController.CreateUser)
	r.POST("/users/:id/bets", usersController.CreateUserBet)

	r.DELETE("/users/:id", usersController.DeleteUser)

	// Register routes
	http.Handle("/", r)
}

func Auth(handler httprouter.Handle, scopesNeeded []string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("authorization")
		if token == "" {
			SharedResponses.Unauthorized(w, errors.New("No token provided"))
			return
		}

		userScopes, err := userService.Authenticate(token)
		if err != nil {
			SharedResponses.Unauthorized(w, err)
			return
		}

		// Validate token comes with scope(s) needed for the requested resource
		for _, scope := range scopesNeeded {
			hasScope := false
			for _, userScope := range userScopes {
				if userScope == scope {
					hasScope = true
				}
			}
			if !hasScope {
				SharedResponses.Unauthorized(w, nil)
				return
			}
		}
		handler(w, r, ps)
	}

}

func startServer() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
