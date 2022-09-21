package router

import (
	"user-service/controller"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine) {
	userService := service.NewUserService()
	notifyService := service.NewNotifyService()
	registrationController := controller.NewRegistrationController(userService, notifyService)
	engine.GET("/users", registrationController.GetUser)
	engine.GET("/user/:id", registrationController.GetUserById)
	engine.POST("/user", registrationController.PostUser)
}
