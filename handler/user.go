package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	formatter := user.FormatUser(newUser, "token")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}
