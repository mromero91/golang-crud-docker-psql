package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mromero91/go-api-contacts/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUser)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
