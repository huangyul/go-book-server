# 基础镜像
FROM ubuntu:20.04

# 复制到镜像里
COPY webook /app/webook

# 设定工作目录
WORKDIR /app

# 入口执行
ENTRYPOINT ["/app/webook"]