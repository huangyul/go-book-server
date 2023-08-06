package web

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/edit", u.Edit)
	ug.POST("/login", u.Login)
	ug.POST("/signup", u.SignUp)
}

// Profile 获取用户信息
func (u *UserHandler) Profile(ctx *gin.Context) {}

// SignUp 注册
func (u *UserHandler) SignUp(ctx *gin.Context) {}

// Login 登录
func (u *UserHandler) Login(ctx *gin.Context) {}

// Edit 编辑
func (u *UserHandler) Edit(ctx *gin.Context) {}
