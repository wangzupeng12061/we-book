package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangzupeng12061/we-book/internal/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHandler{}
	server.POST("/users/login", u.Login)
	server.POST("/users/signup", u.SignUp)
	server.POST("/users/edit", u.Edit)
	server.GET("/users/profile", u.Profile)
	server.Run(":8080")
}
