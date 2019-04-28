package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	mongo "github.com/alindenberg/know-it-all/database"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	matchesRepository "github.com/alindenberg/know-it-all/domain/matches/repository"
	matchesController "github.com/alindenberg/know-it-all/domain/matches/controller"
	leaguesController "github.com/alindenberg/know-it-all/domain/leagues/controller"
)

func main() {
	mongo.InitDatabase()
	addRouteHandlers()
	log.Println("Started Go server")
	createGraphQlSchema()
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

func createGraphQlSchema() {
	var MatchType = graphql.NewObject(graphql.ObjectConfig{
	  Name: "Address",
	  Fields: graphql.Fields{
	    "MatchID": &graphql.Field{ Type: graphql.String },
	    "HomeTeam": &graphql.Field{ Type: graphql.String },
	    "AwayTeam": &graphql.Field{ Type: graphql.String },
	    "Date": &graphql.Field{ Type: graphql.String },
	    "Winner": &graphql.Field{ Type: graphql.Int },
	  },
	});

	var CreateResponse = graphql.NewObject(graphql.ObjectConfig{
		Name: "CreateResponse",
		Fields: graphql.Fields{
			"ID": &graphql.Field{ Type: graphql.String },
		},
	})
	var queryType = graphql.NewObject(graphql.ObjectConfig{
	    Name: "Query",
	    Fields: graphql.Fields{
	        "getMatch": &graphql.Field{
	            Type: MatchType,
	            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	                id := params.Args["id"].(string)
	                return matchesRepository.GetMatch(id), nil
	            },
	            Args: graphql.FieldConfigArgument{
	            	"id": {
	            		Type: graphql.String,
	            		Description: "League ID",
	            	},
	            },
	        },
	        "getMatches": &graphql.Field{
	            Type: graphql.NewList(MatchType),
	            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	                // return matchesController.GetAllMatches()
	                // id := params.Args["id"].(string)
	                allMatches := matchesRepository.GetAllMatches()
	                log.Println(allMatches)
	                return allMatches, nil
	            },
	        },
	    },
	})

	var mutationType = graphql.NewObject(graphql.ObjectConfig{
	    Name: "Mutation",
	    Fields: graphql.Fields{
	        "createMatch": &graphql.Field{
	            Type: CreateResponse,
	            Resolve: func(params graphql.ResolveParams) (interface{}, error) {
	                id := uuid.New().String()
	                homeTeam := params.Args["homeTeam"].(string)
	                awayTeam := params.Args["awayTeam"].(string)
	                date := params.Args["date"].(string)
	                winner := params.Args["winner"].(int)

	                match := MatchModels.Match{id, homeTeam, awayTeam, date, winner}
	                result := matchesRepository.CreateMatch(match)

	                return MatchModels.CreateResponse{result}, nil
	            },
	            Args: graphql.FieldConfigArgument{
	            	"homeTeam": {
	            		Type: graphql.String,
	            		Description: "Home Team",
	            	},
	            	"awayTeam": {
	            		Type: graphql.String,
	            		Description: "Away Team",
	            	},
	            	"date": {
	            		Type: graphql.String,
	            		Description: "Date of the match",
	            	},
	            	"winner": {
	            		Type: graphql.Int,
	            		Description: "Winner of the match",
	            	},
	            },
	        },
	    },
	})

	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	    Query: queryType,
	    Mutation: mutationType,
	})

	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
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
