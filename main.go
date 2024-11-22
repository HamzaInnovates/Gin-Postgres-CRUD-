package main

import (
	"postgres/config"
	"postgres/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	root := gin.New()
	router.UserRoute(root)
	root.Run(":8000")
}
