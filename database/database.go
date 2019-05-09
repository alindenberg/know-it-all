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
	db *mongo.Database
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

	db = client.Database("knowitall")
}

func GetDbClient() *mongo.Database {
	return db
}
