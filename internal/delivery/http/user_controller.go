package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Login(ctx *gin.Context) {
	fmt.Println("Hello World")
	// Handle the request
	data := map[string]string{
		"message": "Hello World",
	}
	ctx.JSON(http.StatusOK, data)
}
