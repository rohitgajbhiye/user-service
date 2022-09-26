package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type LifeCycleController interface {
	PostStart(ctx *gin.Context)
	PreStop(ctx *gin.Context)
}

type lifeCycleController struct {
}

func NewLifeCycleController() LifeCycleController {
	return &lifeCycleController{}
}

func (c *lifeCycleController) PostStart(ctx *gin.Context) {
	log.Println("Hitting post start lifecycle hook")
}

func (c *lifeCycleController) PreStop(ctx *gin.Context) {
	log.Println("Hitting pre stop lifecycle hook")
}
