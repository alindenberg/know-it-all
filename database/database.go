package database

import (
	"context"
	"log"

	// "net/http"
	"time"
	// "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	Db *mongo.Database
)

func InitDatabase() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	Db = client.Database("matches")

	// collection := client.Collection("cities")
	// result := struct{
	//   	Name string
	// 		State string
	// }{}
	// err = collection.FindOne(context.TODO(), bson.D{{"name", "Gary"}}).Decode(&result)
	// if(err != nil) { log.Fatal(err) }
	// log.Println("Find result name : ", result.Name)
	// log.Println("Find result state : ", result.State)
}
