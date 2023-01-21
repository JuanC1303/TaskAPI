package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"taskAPI/db"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// swagger:model
type user struct {

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

var users = []user{}

func main() {
	users = db.Mongodb()
	fmt.Println(users)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://JuanC13:1303@clustertaskapi.ckezybd.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:1303")
}

func getUsers(c *gin.Context) {
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

func postUser(c *gin.Context) {
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
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
}

func getUserByID(c *gin.Context) {
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

func updateUser(c *gin.Context) {
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
	var upuser user
	id := c.Param("id")
	upuser.ID = id
	if err := c.BindJSON(&upuser); err != nil {
		return
	}
	for a, j := range users {
		if j.ID == id {
			c.IndentedJSON(http.StatusOK, j)
			users[a] = upuser
		}

	}

}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
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
	for a, j := range users {
		if j.ID == id {
			c.IndentedJSON(http.StatusOK, j)
			users = RemoveIndex(users, a)
		}

	}

}

func RemoveIndex(s []user, index int) []user {
	return append(s[:index], s[index+1:]...)
}
