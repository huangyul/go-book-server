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

### 使用dockerfile打包成一个镜像

定义dockerfile文件

```dockerfile
# 基础镜像
FROM ubuntu:20.04

# 复制到镜像里
COPY webook /app

# 设定工作目录
WORKDIR /app

# 入口执行
ENTRYPOINT ["/app/webook"]
```

执行构建镜像命令

`docker build -t jojo/webook:v0.0.1 .`

`-t`: 镜像的名字及标签，通常 name:tag 或者 name 格式

### 启动k8s

定义k8s文件

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webook
spec:
  replicas: 3
  selector:
    matchLabels:
      app: webook
  # 描述的是POD是什么样的
  template:
    metadata:
      labels:
        app: webook
    # POD的具体信息
    spec:
      containers:
        - name: webook
          image: jojo/webook:v0.0.1
```

启动

`kubectl.exe apply -f .\k8s-webook-deployment.yaml`

查看是否启动成功

`kubectl.exe get deployments`

### k8s持久化mysql

在`deployment`中声明挂载的`volume`，通过`claimName`去匹配对应的`pvc`，然后在`pvc`通过`storageClassName`匹配`pv`

### k8s部署redis

### k8s部署nginx

- Ingress 是路由规则
- Ingress controller控制这些规则，执行这些配置

类似与`Ingress`是配置，`Ingress controller`是执行器

## 性能测试

wrk

## 需求分析

如何快速分析需求：

1. 参考竞品
2. 从不同角度分析
    1. 功能角度
    2. 非功能角度：安全性，拓展性，性能
3. 从正常和异常流程两种角度思考

### 服务划分

短信服务->验证码->验证码登录