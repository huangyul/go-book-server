package web

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"go-book-server/internal/domain"
	"go-book-server/internal/service"
	"net/http"
)

type UserHandler struct {
	svc      *service.UserService
	emailExp *regexp.Regexp
	password *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const emailReg = `^[\w.-]+@[a-zA-Z\d.-]+\.[a-zA-Z]{2,}$`
	const passwordReg = ""
	emailExp := regexp.MustCompile(emailReg, regexp.None)
	passwordExp := regexp.MustCompile(passwordReg, regexp.None)
	return &UserHandler{
		emailExp: emailExp,
		password: passwordExp,
		svc:      svc,
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
	// 获取数据
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
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一样")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "邮箱格式不正确")
		return
	}

	// 处理数据
	err = u.svc.SignUp(ctx, domain.User{Email: req.Email, Password: req.Password})

	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusBadRequest, "邮箱冲突")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	// 返回响应
	ctx.String(http.StatusOK, "注册成功")
}

// Login 登录
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// 查找用户
	du := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := u.svc.FindByEmail(ctx, du)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	// 登录成功
	ctx.String(http.StatusOK, "登录成功")
	return
}

// Edit 编辑
func (u *UserHandler) Edit(ctx *gin.Context) {}
