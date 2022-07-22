package main

import (
	"login-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routes.InitRoutesUser(route)
	route.Run(":8080")
}
