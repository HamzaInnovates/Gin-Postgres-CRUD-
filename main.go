package main

import (
	"postgres/config"
	"postgres/controller"
	"postgres/initializers"
	"postgres/router"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	config.Connect()
	root := gin.New()
	userController := controller.Newusercontroller()
	router.UserRoute(root, userController)
	root.Run(":8000")
}
