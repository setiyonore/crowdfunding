package handler

import (
	"crowdfunding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userhandler {
	return &userhandler{userService}
}

func (h *userhandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userhandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *userhandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_new.html", input)
		return
	}

	registerInput := user.RegisterUserInput{}
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusFound, "/users")

}

func (h *userhandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	registeredUser, err := h.userService.GetUserById(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	input := user.FormUpdateUserInput{}
	input.ID = registeredUser.ID
	input.Name = registeredUser.Name
	input.Email = registeredUser.Email
	input.Occupation = registeredUser.Occupation

	c.HTML(http.StatusOK, "user_edit.html", input)
}

func (h *userhandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input user.FormUpdateUserInput
	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}
	input.ID = id
	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.Redirect(http.StatusFound, "/users")
}
