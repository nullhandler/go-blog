package main

import (
	"go_crud/controllers"
	"go_crud/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Log queries
	models.DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	r.GET("/posts", controllers.FindPosts)

	r.GET("/posts/:id", controllers.FindPost)

	r.POST("/create_post", controllers.CreatePost)

	r.POST("/comment", controllers.CreateComment)

	r.Run("localhost:8080")
}
