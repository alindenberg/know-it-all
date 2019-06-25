package main

import (
	"errors"
	"log"
	"net/http"

	mongo "github.com/alindenberg/know-it-all/database"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
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
	// r.GET("/matches/:id", matchesController.GetMatch)
	// r.GET("/matches", Auth(matchesController.GetAllMatches, []string{"read:matches"}))
	// r.POST("/matches", Auth(matchesController.CreateMatch, []string{"write:matches"}))
	// r.POST("/matches/:id/resolve", Auth(matchesController.ResolveMatch, []string{"update:matches"}))
	// r.DELETE("/matches/:id", Auth(matchesController.DeleteMatch, []string{"delete:matches"}))

	// League Routes
	r.GET("/leagues/:id", Auth(leaguesController.GetLeague, []string{"read:leagues"}))
	r.POST("/leagues/:leagueId/matches", Auth(leaguesController.CreateLeagueMatch, []string{"write:matches"}))
	// r.POST("/matches/:id/resolve", Auth(leaguesController.ResolveMatch, []string{"update:matches"}))
	r.GET("/leagues", leaguesController.GetAllLeagues)
	r.POST("/leagues", leaguesController.CreateLeague)
	r.DELETE("/leagues/:id", leaguesController.DeleteLeague)

	// User Routes
	r.GET("/users/:id", usersController.GetUser)
	r.GET("/users", usersController.GetAllUsers)
	r.POST("/users", usersController.CreateUser)
	r.POST("/users/:id/bets", Auth(usersController.CreateUserBet, []string{"write:bets"}))
	r.DELETE("/users/:id", usersController.DeleteUser)

	// // Bet routes
	// r.GET("/bets", Auth(matchesController.GetAllBets, []string{"read:bets"}))
	// r.POST("/bets", matchesController.CreateBet)
	// r.DELETE("/bets/:betId", matchesController.DeleteBet)

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
