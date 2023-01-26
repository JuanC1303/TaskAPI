package databs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// swagger:model
type User struct {
	// the id for this user
	//required: true
	// min length: 1
	ID string `json:"id"`

	// the name for this user
	//required: true
	// min length: 3
	Name string `json:"name"`

	// the lastname for this user
	// required: true
	// min length: 3
	Lastname string `json:"lastname"`

	// the name for this user
	// required: true
	// Minimum: 21
	Age int `json:"age"`
}

func Mongodb() []User {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	var users = []User{}

	usersCollection := client.Database("usersdb").Collection("users")
	collection, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	} else {
		for collection.Next(ctx) {
			var result User
			err := collection.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				users = append(users, result)
			}
		}
	}
	return users
}
