package matchrepository

import (
	"context"

	mongo "github.com/alindenberg/know-it-all/database"
	MatchModels "github.com/alindenberg/know-it-all/domain/matches/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION = "matches"

func GetAllMatches(leagueID string) ([]*MatchModels.Match, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)

	filter := bson.D{}
	if leagueID != "" {
		filter = bson.D{{"leagueid", leagueID}}
	}

	cur, err := collection.Find(context.TODO(), filter, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*MatchModels.Match
	for cur.Next(context.TODO()) {
		var elem MatchModels.Match
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetMatch(matchID string) (*MatchModels.Match, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	match := MatchModels.Match{}
	err := collection.FindOne(context.TODO(), bson.D{{"matchid", matchID}}).Decode(&match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}

func CreateMatch(match *MatchModels.Match) error {
	_, err := mongo.GetDbClient().Collection(COLLECTION).InsertOne(context.TODO(), &match)
	if err != nil {
		return err
	}
	return nil
}

func ResolveMatch(matchID string, matchResult *MatchModels.MatchResult) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.D{
			{"matchid", matchID},
		},
		bson.D{
			{"$set", bson.D{
				{"awayteamscore", matchResult.AwayScore},
				{"hometeamscore", matchResult.HomeScore},
				{"isresolved", true},
			}},
		},
	)

	return err
}
