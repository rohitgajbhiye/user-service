package controller

import (
	"fmt"
	"net/http"
	"user-service/controller/mapper"
	"user-service/controller/request"
	"user-service/controller/response"
	"user-service/model"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

type RegistrationController interface {
	GetUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	PostUser(ctx *gin.Context)
}

type registrationController struct {
	userService   service.UserService
	notifyService service.NotifyService
}

func NewRegistrationController(userService service.UserService, notifyService service.NotifyService) RegistrationController {
	return &registrationController{
		userService:   userService,
		notifyService: notifyService,
	}
}

func (c *registrationController) GetUser(ctx *gin.Context) {
	users, err := c.userService.FetchAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			response.Error{
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Not user found",
				ErrorType:    "Internal",
			})
		return
	}
	var userArr []model.User
	for _, v := range users {
		userArr = append(userArr, v)
	}
	ctx.JSON(http.StatusOK, userArr)
	return
}

func (c *registrationController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest,
			response.Error{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "ID not found",
				ErrorType:    "Validation",
			})
		return
	}
	user, err := c.userService.FetchUserById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			response.Error{
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Not user found",
				ErrorType:    "Internal",
			})
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

func (c *registrationController) PostUser(ctx *gin.Context) {
	var request request.RegistrationRequest
	valErr := ctx.BindJSON(&request)
	if valErr != nil {
		ctx.JSON(http.StatusBadRequest,
			response.Error{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: valErr.Error(),
				ErrorType:    "Validation",
			})
		return
	}
	id, err := c.userService.RegisterUser(ctx,
		mapper.RegistrationToUserMapper(request))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			response.Error{
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "Error creating user",
				ErrorType:    "Internal",
			})
	}
	// Notify to notification service
	notifyErr := c.notifyService.Notify(ctx, id, "ACCOUNT_CREATED")
	if notifyErr != nil {
		ctx.JSON(http.StatusInternalServerError,
			response.Error{
				ErrorCode:    http.StatusInternalServerError,
				ErrorMessage: "failed to send notification",
				ErrorType:    "Internal",
			})
		fmt.Println(notifyErr)
		return
	}
	ctx.Status(http.StatusCreated)
}
