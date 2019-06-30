package userrepository

import (
	"context"
	"strings"

	mongo "github.com/alindenberg/know-it-all/database"
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var SORT = bson.D{{"winpercentage", -1}}

func GetLeaderboard() ([]*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(SORT))
	if err != nil {
		return nil, err
	}

	results := []*UserModels.User{}
	for cur.Next(context.TODO()) {
		var entry UserModels.User
		err := cur.Decode(&entry)
		if err != nil {
			return nil, err
		}
		results = append(results, &entry)
	}

	cur.Close(context.TODO())

	return results, nil
}

// Get a leaderboard for a user - contains user and friends
func GetLeaderboardOnlyForIds(userIds []string) ([]*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	results := []*UserModels.User{}

	cur, err := collection.Find(
		context.TODO(),
		bson.D{
			{"userid", bson.D{{"$in", userIds}}},
		},
		options.Find().SetSort(SORT),
	)
	if err != nil {
		// log.Println(err.Error())
		if strings.Contains(err.Error(), "no documents in result") {
			return results, nil
		}
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var entry UserModels.User
		err := cur.Decode(&entry)
		if err != nil {
			return nil, err
		}
		results = append(results, &entry)
	}

	cur.Close(context.TODO())

	return results, nil
}
