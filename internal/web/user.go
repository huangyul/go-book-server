package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	emailExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const emailReg = `/^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,}$/
`
	emailExp := regexp.MustCompile(emailReg, regexp.None)
	return &UserHandler{
		emailExp: emailExp,
	}
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
func (u *UserHandler) SignUp(ctx *gin.Context) {
	// 1 获取数据
	type Req struct {
		Email           string `json:"email,omitempty"`
		Password        string `json:"password,omitempty"`
		ConfirmPassword string `json:"confirm_password,omitempty"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		// 这里gin会自动返回400
		return
	}

	// 校验数据
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		// 这里是正则表达式不对
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱格式不正确")
		return
	}

	// 处理数据

	// 返回响应
	ctx.String(http.StatusOK, "注册成功")
}

// Login 登录
func (u *UserHandler) Login(ctx *gin.Context) {}

// Edit 编辑
func (u *UserHandler) Edit(ctx *gin.Context) {}
