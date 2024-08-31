package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func (u UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/edit", u.Edit)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.SignUp)
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
