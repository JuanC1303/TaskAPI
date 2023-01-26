package main

import (
	gu "github.com/JuanC1303/TaskAPI/apiHandler"
	"github.com/JuanC1303/TaskAPI/config"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/users", gu.Getusers)
	router.GET("/users/:id", gu.GetUserByID)
	router.POST("/users", gu.PostUser)
	router.PUT("/users/:id", gu.UpdateUser)
	router.DELETE("/users/:id", gu.DeleteUser)

	router.Run(config.ROUTER)
}
