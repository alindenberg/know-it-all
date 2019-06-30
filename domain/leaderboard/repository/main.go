package leaderboardrepository

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION := "leaderboard"

func GetLeaderboard() ([]*LeaderboardModels.LeaderboardEntry, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*LeaderboardModels.LeaderboardEntry
	for cur.Next(context.TODO()) {
		var entry LeaderboardModels.LeaderboardEntry
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
func GetUserFriendsLeaderboard(userIds []string) ([]*LeaderboardModels.LeaderboardEntry, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{{"userid", {"$in", userIds}}}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*LeaderboardModels.LeaderboardEntry
	for cur.Next(context.TODO()) {
		var entry LeaderboardModels.LeaderboardEntry
		err := cur.Decode(&entry)
		if err != nil {
			return nil, err
		}
		results = append(results, &entry)
	}

	cur.Close(context.TODO())

	return results, nil
}
