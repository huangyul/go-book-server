package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-book-server/internal/web"
	"strings"
	"time"
)

func main() {
	server := gin.Default()

	// 解决跨域问题
	server.Use(cors.New(cors.Config{
		AllowCredentials: true, // 是否允许使用cookie
		AllowHeaders:     []string{"Content-Type", "Authed"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "your-company.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	// 注册用户路由
	user := web.NewUserHandler()
	user.RegisterRoutes(server)

	err := server.Run(":8888")
	if err != nil {
		return
	}
}
