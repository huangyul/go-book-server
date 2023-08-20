package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go-book-server/internal/repository"
	"go-book-server/internal/repository/dao"
	"go-book-server/internal/service"
	"go-book-server/internal/web"
	"go-book-server/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()

	server := initWeb()

	user := initUser(db)

	// 注册用户路由
	user.RegisterRoutes(server)

	err := server.Run(":8888")
	if err != nil {
		return
	}
}

func initWeb() *gin.Engine {
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

	// 设置session
	// 这里是使用cookie存储session
	//store := cookie.NewStore([]byte("secret"))
	// 使用redis存储,便于多实例
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("pO8uY2tE6iK9zN4oT2wP3iP4mQ9vN6fC"), []byte("oU5yR0iC6zM9uF0gR3vX0iR0qL1hR1zA"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))

	// session校验
	server.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/users/signup").IgnorePaths("/users/login").Build())

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	user := web.NewUserHandler(svc)

	err := dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return user
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	return db
}
