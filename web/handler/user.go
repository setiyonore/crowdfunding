package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
}

func NewUserHandler() *userhandler {
	return &userhandler{}
}

func (h *userhandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "user_index.html", nil)
}
