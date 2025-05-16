# 第一阶段：构建 Go 应用
FROM golang:1.22 AS builder

# 设置国内 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# 下载依赖并构建（关闭 CGO）
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o wechat ./main.go

# 第二阶段：运行时镜像（使用更小的基础镜像）
FROM debian:bullseye-slim

WORKDIR /app

# 安装运行所需的最基本工具（如有需要）
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# 拷贝构建后的 Go 可执行文件及配置
COPY --from=builder /app/wechat .
COPY config/app.yml /app/config/app.yml

# 设置暴露端口
EXPOSE 8081

# 启动命令
CMD ["/bin/sh", "-c", "/app/wechat migrate up && exec /app/wechat"]
