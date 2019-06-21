package main

import (
	"errors"
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
	r.GET("/matches", Auth(matchesController.GetAllMatches, []string{"read:matches"}))
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
	r.GET("/bets", Auth(matchesController.GetAllBets, []string{"read:bets"}))
	r.POST("/bets", matchesController.CreateBet)
	r.DELETE("/bets/:betId", matchesController.DeleteBet)

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
