package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type LoginJWTMiddlewareBuilder struct {
	// 不需要进行登录校验的地址
	paths []string
}

func NewLoginJWTMiddlewareBuild() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 不需要登录校验的
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		// 拿到header的token进行校验
		token := ctx.GetHeader("Authorization")
		if token == "" {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		segs := strings.Split(token, " ")
		if len(segs) != 2 {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := segs[1]
		t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if t == nil || !t.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
