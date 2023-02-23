package gu

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	mdb "workspace/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

var usersCollection *mongo.Collection = mdb.Conectmongo()

func Getusers(c *gin.Context) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
				c.IndentedJSON(http.StatusOK, result)
			}
		}
	}

}

func GetUserByID(c *gin.Context) {
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
	var user User
	error := usersCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, "Not found")
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func PostUser(c *gin.Context) {

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
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Check the information you want to add.")
	} else {
		newUser.ID = (uuid.New()).String()
		res, err := usersCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "error")
		}
		fmt.Println(res)
		c.IndentedJSON(http.StatusOK, newUser.ID)
	}
}

func DeleteUser(c *gin.Context) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Param("id")
	res, err := usersCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error")
	}
	if res.DeletedCount == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Not found")
	} else {
		c.IndentedJSON(http.StatusOK, res)
	}
}

func UpdateUser(c *gin.Context) {
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
	var upuser User
	if err := c.BindJSON(&upuser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Check the information you want to add.")
	} else {
		id := c.Param("id")
		newUser := User{ID: id, Name: upuser.Name, Lastname: upuser.Lastname, Age: upuser.Age}
		res, err := usersCollection.ReplaceOne(context.TODO(), bson.M{"id": id}, newUser)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Error")
		}
		if res.MatchedCount == 0 {
			c.IndentedJSON(http.StatusBadRequest, "Not found")
		} else {
			c.IndentedJSON(http.StatusOK, newUser)
		}
	}
}
