package pu

import (
	"C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/databs"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func PostUser(c *gin.Context) {
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
	// swagger:operation POST /users postUser
	//
	// post a new user
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
	var newUser databs.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser.ID = (uuid.New()).String()
	res, err := usersCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
