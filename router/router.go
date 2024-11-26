package router

import (
	"postgres/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, usercontroller controller.UserController) {
	router.GET("/", usercontroller.GetUserData)
	router.POST("/user", usercontroller.AddUserData)
	router.PUT("/user/:id", usercontroller.UpdateUserData)
	router.DELETE("/user/:id", usercontroller.DeleteUserData)
	router.GET("/user/:id", usercontroller.GetUser)
	router.POST("/signin", usercontroller.AuthenticateUser)

}
