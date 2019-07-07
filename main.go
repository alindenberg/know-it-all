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
	r.GET("/leagues/:leagueId", leaguesController.GetLeague)
	r.GET("/leagues/:leagueId/matches", leaguesController.GetAllLeagueMatches)
	r.GET("/leagues/:leagueId/matches/:matchId", leaguesController.GetLeagueMatch)
	r.GET("/leagues/:leagueId/teams", teamsController.GetAllTeamsForLeague)

	r.POST("/leagues", leaguesController.CreateLeague)
	r.POST("/leagues/:leagueId/matches", leaguesController.CreateLeagueMatch)
	r.POST("/leagues/:leagueId/matches/:matchId/resolve", leaguesController.ResolveLeagueMatch)

	r.DELETE("/leagues/:id", leaguesController.DeleteLeague)

	// User Routes
	r.GET("/users", usersController.GetAllUsers)
	r.GET("/users/:userId", usersController.GetUser)
	r.GET("/users/:userId/friends", usersController.GetUserFriends)

	r.POST("/users", usersController.CreateUser)
	r.POST("/users/:userId/username", usersController.CreateUsername)
	r.POST("/users/:userId/bets", usersController.CreateUserBet)
	r.POST("/users/:userId/bets/:matchId", usersController.UpdateUserBet)
	r.POST("/users/:userId/friends/:friendId", usersController.AddFriend)

	r.DELETE("/users/:userId", usersController.DeleteUser)
	r.DELETE("/users/:userId/friends/:friendId", usersController.DeleteUserFriend)

	// Leaderboard routes
	r.GET("/leaderboard", usersController.GetLeaderboard)
	r.GET("/leaderboard/:userId", usersController.GetLeaderboardForUser)

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
