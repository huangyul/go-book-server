package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-book-server/internal/repository"
	"go-book-server/internal/repository/dao"
	"go-book-server/internal/service"
	"go-book-server/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	server := gin.Default()

	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	user := web.NewUserHandler(svc)

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
	user.RegisterRoutes(server)

	err = server.Run(":8888")
	if err != nil {
		return
	}
}
