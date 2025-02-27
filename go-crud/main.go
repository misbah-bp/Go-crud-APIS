package main

import (
	"github.com/gin-gonic/gin"
	"github.com/misbah-bp/go-crud/controllers"
	"github.com/misbah-bp/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/blog", controllers.BlogCreate)
	r.GET("/blog", controllers.BlogGet)
	r.GET("/blog/:id", controllers.SingelBlog)
	r.PUT("/blog/:id", controllers.BlogUpdate)
	r.DELETE("/blog/:id", controllers.DeleteBlog)

	r.Run()
}
