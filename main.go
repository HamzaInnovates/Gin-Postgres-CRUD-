package main

import (
	"postgres/config"
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
	router.UserRoute(root)
	root.Run(":8000")
}
