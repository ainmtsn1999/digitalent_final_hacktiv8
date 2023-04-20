package routes

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/handlers"
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApi(router *gin.Engine, handler handlers.HttpServer) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", handler.GetAllPhoto)
		photoRouter.GET("/:photoId", handler.GetPhoto)
		photoRouter.POST("/", handler.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuth(), handler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuth(), handler.DeletePhoto)
		photoRouter.POST("/:photoId/comment", handler.CreateComment)

	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", handler.GetAllComment)
		commentRouter.GET("/:commentId", handler.GetComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuth(), handler.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuth(), handler.DeleteComment)

	}
}
