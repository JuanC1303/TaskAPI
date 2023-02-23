package main

import (
	gu "workspace/api"
	mdb "workspace/mongodb"
	"workspace/tokens"

	"github.com/gin-gonic/gin"
)

func main() {

	// http.Handle("/api", ValidateJWT(gu.Getusers))
	// http.HandleFunc("/jwt", GetJwt)

	// http.ListenAndServe(":3500", nil)
	router := gin.Default()
	router.GET("/token", tokens.GetJwt)
	secured := router.Group("/secured").Use(tokens.Auth())
	{
		secured.GET("/users", gu.Getusers)
		secured.GET("/users/:id", gu.GetUserByID)
		secured.POST("/users", gu.PostUser)
		secured.PUT("/users/:id", gu.UpdateUser)
		secured.DELETE("/users/:id", gu.DeleteUser)
	}

	router.Run(mdb.Config().GetString("ROUTER"))
}
