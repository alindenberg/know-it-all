package userrepository

import (
	"context"
	"errors"
	"fmt"
	"log"

	mongo "github.com/alindenberg/know-it-all/database"
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var COLLECTION = "users"

func GetAllUsers() ([]*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*UserModels.User
	for cur.Next(context.TODO()) {
		var user UserModels.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		results = append(results, &user)
	}

	cur.Close(context.TODO())

	return results, nil
}

func GetUser(id string) (*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result := UserModels.User{}
	err := collection.FindOne(context.TODO(), bson.D{{"userid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func GetUserByUsername(username string) (*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result := UserModels.User{}
	err := collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateUser(user *UserModels.User) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), &user)

	return err
}

func DeleteUser(id string) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"userid", id}})

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(fmt.Sprintf("Document with id %s was not found", id))
	}
	return nil
}

func CreateUserBet(id string, bet *UserModels.UserBet) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	res, err := collection.UpdateOne(
		context.TODO(),
		bson.D{
			{"userid", id}, {"bets.matchid", bet.MatchID},
		},
		bson.D{
			{"$set", bson.D{{"bets.$", bet}}},
		},
	)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		res, err = collection.UpdateOne(
			context.TODO(),
			bson.D{
				{"userid", id},
			},
			bson.D{
				{"$push", bson.D{{"bets", bet}}},
			},
		)
	}
	return nil
}

func AddFriend(userId string, friendId string) error {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.D{{"userid", userId}},
		bson.D{{"$push", bson.D{{"friends", friendId}}}},
	)
	return err
}

func GetUsersWithBetOnMatch(matchID string) ([]*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(
		context.TODO(),
		bson.D{
			{"bets.matchid", matchID},
		},
	)
	if err != nil {
		return nil, err
	}

	var results []*UserModels.User
	for cur.Next(context.TODO()) {
		var user UserModels.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		results = append(results, &user)
	}

	cur.Close(context.TODO())

	return results, err
}

func UpdateUser(user *UserModels.User) {
	log.Println("USER IN REPO ", user)
	collection := mongo.GetDbClient().Collection(COLLECTION)
	res, err := collection.ReplaceOne(
		context.TODO(),
		bson.D{
			{"userid", user.UserID},
		},
		user,
	)

	log.Println("update user err : ", err)
	log.Println("update user result : ", res)
}
