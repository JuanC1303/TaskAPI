package du

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
