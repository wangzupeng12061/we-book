package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func (u UserHandler) RegisterRoutes(server *gin.Engine) {
	server.POST("/users/login", u.Login)
	server.POST("/users/signup", u.SignUp)
	server.POST("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)
}
func (u *UserHandler) SignUp(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, signup")

}
func (u UserHandler) Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, login")

}
func (u UserHandler) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, edit")
}
func (u UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, profile")

}
