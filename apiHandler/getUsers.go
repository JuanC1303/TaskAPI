package gu

import (
	databs "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/db"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Getusers(c *gin.Context) {
	users := databs.Mongodb()
	// swagger:operation GET /users getUsers
	//
	// Returns all users
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: user response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/user"
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	users := databs.Mongodb()
	// swagger:operation GET /users/:id getUsersbyID
	//
	// Returns one user define by his id
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: user response
	//     schema:
	//       type: array
	//       items:
	//         "$ref": "#/definitions/user"
	id := c.Param("id")

	for _, j := range users {
		if j.ID == id {
			c.IndentedJSON(http.StatusOK, j)
			return
		}
	}
}

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

func DeleteUser(c *gin.Context) {
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
	// swagger:operation DELETE /users/:id deleteUser
	//
	// delete one user define by his id
	//
	// ---
	// produces:
	// - application/json
	// responses:
	//   '200':
	//     description: user response
	usersCollection := client.Database("usersdb").Collection("users")
	id := c.Param("id")
	users := databs.Mongodb()
	for _, j := range users {
		if j.ID == id {
			res, err := usersCollection.DeleteOne(ctx, bson.M{"id": j.ID})
			if err != nil {
				log.Fatal(err)
			}
			c.IndentedJSON(http.StatusOK, res)
		}
	}
}

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
