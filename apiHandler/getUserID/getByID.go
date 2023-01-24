package guid

import (
	"C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/databs"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
