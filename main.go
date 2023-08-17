package main

import (
	"github.com/Shaheer_25/controllers" // Corrected import path
	"github.com/Shaheer_25/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitializeVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostShow)
	r.GET("/posts/:id", controllers.SinglePostShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}

