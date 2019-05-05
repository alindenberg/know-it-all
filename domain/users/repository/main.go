package userrepository

import (
	"fmt"
	"errors"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	mongo "github.com/alindenberg/know-it-all/database"
	UserModels "github.com/alindenberg/know-it-all/domain/users/models"
)

var COLLECTION = "users"

func GetAllUsers() ([]*UserModels.User, error) {
	collection := mongo.Db.Collection(COLLECTION)
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
	collection := mongo.Db.Collection(COLLECTION)
	result := UserModels.User{}
	err := collection.FindOne(context.TODO(), bson.D{{"userid", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreateUser(user UserModels.User) error {
	collection := mongo.Db.Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), user)

	return err
}

func DeleteUser(id string) error {
	collection := mongo.Db.Collection(COLLECTION)
	result, err := collection.DeleteOne(context.TODO(), bson.D{{"userid", id}})

	if err != nil {
		return err
	}

	if(result.DeletedCount == 0) {
		return errors.New(fmt.Sprintf("Document with id %s was not found", id))
	}
	return nil
}
