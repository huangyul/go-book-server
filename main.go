package main

import (
	"github.com/gin-gonic/gin"
	"go-book-server/internal/web"
)

func main() {
	server := gin.Default()
	user := web.UserHandler{}
	user.RegisterRoutes(server)

	err := server.Run(":8888")
	if err != nil {
		return
	}
}
