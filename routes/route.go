package routes

import (
	"github.com/ainmtsn1999/digitalent_final_hacktiv8/handlers"
	"github.com/gin-gonic/gin"
)

func InitApi(router *gin.Engine, handler handlers.HttpServer) {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.Register)
		userRouter.POST("/login", handler.Login)
	}

}
