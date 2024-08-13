package main

import (
	"github.com/abhi1510/go-rest-api/controllers"
	"github.com/abhi1510/go-rest-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.GetAllPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
