package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	// 告诉golang使用什么编码
	gob.Register(time.Time{})
	return func(ctx *gin.Context) {

		//if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}

		// 判断哪些请求不需要校验
		for _, path := range l.paths {
			if path == ctx.Request.URL.Path {
				return
			}
		}

		sess := sessions.Default(ctx)
		userId := sess.Get("userId")
		if userId == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		const timeKey = "update_time"
		val := sess.Get(timeKey)
		// 判断取出来的数据是不是时间类型
		updateTime, ok := val.(time.Time)
		if val == nil || (ok && time.Now().Sub(updateTime) > time.Second*60) {
			sess.Options(sessions.Options{MaxAge: 60 * 60})
			sess.Set(timeKey, time.Now())
			if err := sess.Save(); err != nil {
				panic(err)
			}
		}
	}
}
