package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangzupeng12061/we-book/internal/web"
)

func main() {
	server := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRoutes(server)
	server.Run(":8080")
}
