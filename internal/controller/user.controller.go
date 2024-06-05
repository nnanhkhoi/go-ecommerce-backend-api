package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nnanhkhoi/go-ecommerce-backend-api/internal/service"
	"github.com/nnanhkhoi/go-ecommerce-backend-api/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {

	response.SuccessResponse(c, 20001, []string{"tipjs", "m10", "cr7"})

	// response.ErrorResponse(c, 20003, "No need")

}
