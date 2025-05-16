# 第一阶段：构建 Go 应用
FROM golang:1.22 AS builder

# 设置国内 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# 下载依赖并构建（关闭 CGO）
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o wechat ./main.go

# 第二阶段：使用最小的运行镜像（Alpine）
FROM alpine:3.19

WORKDIR /app

# 替换 apk 源为阿里云镜像，加速依赖安装
# 加速源 + 安装 ca-certificates + tzdata
RUN sed -i 's|http://dl-cdn.alpinelinux.org|https://mirrors.aliyun.com|g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

# 设置默认时区
ENV TZ=Asia/Shanghai

# 拷贝构建后的 Go 可执行文件及配置
COPY --from=builder /app/wechat .
COPY config/app.yml /app/config/app.yml

# 设置暴露端口
EXPOSE 8082

# 启动命令
CMD ["/bin/sh", "-c", "/app/wechat migrate up && exec /app/wechat"]
