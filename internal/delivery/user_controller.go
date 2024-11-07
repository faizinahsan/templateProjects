package delivery

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
	ctx.String(http.StatusOK, "Hello, World!")
}
