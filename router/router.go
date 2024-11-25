package router

import (
	"postgres/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUserData)
	router.POST("/user", controller.AddUserData)
	router.PUT("/user/:id", controller.UpdateUserData)
	router.DELETE("/user/:id", controller.DeleteUserData)
	router.GET("/user/:id", controller.GetUser)
	router.POST("/signin", controller.AuthenticateUser)

}
