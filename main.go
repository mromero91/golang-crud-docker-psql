package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mromero91/go-api-contacts/config"
	routes "github.com/mromero91/go-api-contacts/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
