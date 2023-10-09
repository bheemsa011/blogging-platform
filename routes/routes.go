// post_routes.go
package routes

import (
	"blogging-platform/handlers"

	"github.com/gin-gonic/gin"
)

// setting up the api group
func SetupPostRoutes(router *gin.RouterGroup) {
	postRoutes := router.Group("/post")
	{
		postRoutes.GET("/:id", handlers.GetPostByID)
		postRoutes.POST("/", handlers.CreatePost)
		postRoutes.PUT("/:id", handlers.UpdatePost)
		postRoutes.DELETE("/:id", handlers.DeletePostByID)
	}
}

func SetupPostsRoutes(router *gin.RouterGroup) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("/", handlers.GetPosts)
	}
}
