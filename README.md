# go server

## 目录接口

- main.go 入口文件
- internal 放业务代码
- pkg 用来放可以被其他项目使用的代码

## 跨域

协议、域名和端口任意一个不同，都是跨域请求

解决方法：后端接收到frefilght请求时，返回允许的请求方法和请求头

使用gin-cros中间件

使用alloworginfunc可以动态配置origin

## gorm

## k8s

### 先将项目打包成linux可执行文件

windows下的命令

```bash
# windows 可执行文件
go evn -w CGO_ENABLED=0
go evn -w GOOS=darwin
go evn -w GOARCH=amd64
go build main.go
# linux 可执行文件
go evn -w CGO_ENABLED=0
go evn -w GOOS=linux
go evn -w GOARCH=amd64
go build main.go
```