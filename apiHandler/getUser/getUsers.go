package gu

import (
	"C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/databs"
	"net/http"

	"github.com/gin-gonic/gin"
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
