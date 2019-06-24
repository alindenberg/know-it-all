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
var USER_KEYS_COLLECTION = "user_keys"

func GetAllUsers() ([]*UserModels.User, error) {
	collection := mongo.GetDbClient().Collection(COLLECTION)
	cur, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var results []*UserModels.User
	for cur.Next(context.TODO()) {
		var elem UserModels.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, &elem)
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
	log.Println("Match id : ", bet.MatchID)
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
