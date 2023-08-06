package main

import (
	"github.com/gin-gonic/gin"
	"go-book-server/internal/web"
)

func main() {
	server := gin.Default()

	// 注册用户路由
	user := web.NewUserHandler()
	user.RegisterRoutes(server)

	err := server.Run(":8888")
	if err != nil {
		return
	}
}
