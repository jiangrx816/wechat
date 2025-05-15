# 第一阶段：构建 Go 应用
FROM golang:1.22 AS builder

# 设置国内 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# 下载依赖并构建（关闭 CGO）
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o tool ./main.go

# 第二阶段：运行时镜像
FROM python:3.10-bullseye

WORKDIR /app

# 替换源（确认 sources.list 存在后再替换）
RUN test -f /etc/apt/sources.list && \
    sed -i 's|http://deb.debian.org|https://mirrors.tuna.tsinghua.edu.cn|g' /etc/apt/sources.list && \
    apt-get update && \
    apt-get install -y netcdf-bin && \
    rm -rf /var/lib/apt/lists/*

# pip 使用阿里源
RUN pip config set global.index-url https://mirrors.aliyun.com/pypi/simple

# 安装依赖
RUN pip install \
    pandas==2.2.2 \
    numpy==1.26.4 \
    geopandas==1.0.1 \
    numpydoc==1.7.0 \
    shapely==2.0.6 \
    openpyxl==3.1.5 \
    xarray==2025.3.1 \
    h5netcdf==1.6.1 \
    tqdm==4.67.1 \
    netCDF4==1.7.2

# 拷贝执行文件
COPY --from=builder /app/tool .
COPY config/app.yml /app/config/app.yml

EXPOSE 8080
CMD ["/bin/sh", "-c", "/app/tool migrate up && exec /app/tool"]