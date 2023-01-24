package main

import (
	du "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/apiHandler/deleteUser"
	gu "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/apiHandler/getUser"
	guid "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/apiHandler/getUserID"
	pu "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/apiHandler/postUser"
	uu "C/Users/juuan/Documents/progra/Golang/taskAPI/TaskAPI/apiHandler/updateUser"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/users", gu.Getusers)
	router.GET("/users/:id", guid.GetUserByID)
	router.POST("/users", pu.PostUser)
	router.PUT("/users/:id", uu.UpdateUser)
	router.DELETE("/users/:id", du.DeleteUser)

	router.Run("localhost:1303")
}
