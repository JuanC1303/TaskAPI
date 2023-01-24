package uu

import (
	"C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/databs"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func UpdateUser(c *gin.Context) {
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
	// swagger:operation PUT /users/:id updateUser
	//
	// update a user that already exist in the array
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: Body
	//   in: body
	//   description: The new user
	//   schema:
	//     "$ref": "#/definitions/user"
	// responses:
	//   '200':
	// 		description: user response
	usersCollection := client.Database("usersdb").Collection("users")
	var upuser databs.User
	if err := c.BindJSON(&upuser); err != nil {
		return
	}
	id := c.Param("id")
	users := databs.Mongodb()
	for _, j := range users {
		if j.ID == id {
			newUser := databs.User{ID: id, Name: upuser.Name, Lastname: upuser.Lastname, Age: upuser.Age}
			res, err := usersCollection.ReplaceOne(context.TODO(), bson.M{"id": j.ID}, newUser)
			if err != nil {
				log.Fatal(err)
			}
			c.IndentedJSON(http.StatusOK, res)
		}
	}
}
