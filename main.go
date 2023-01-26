package main

import (
	gu "https://github.com/JuanC1303/TaskAPI/apiHandler"
	"https://github.com/JuanC1303/TaskAPI/config"

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
