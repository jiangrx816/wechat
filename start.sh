#!/bin/bash

echo "等待 MySQL 启动完成..."

# 等待 mysql 容器的健康检查通过
until docker inspect --format "{{json .State.Health.Status }}" mysql-db | grep -q '"healthy"'; do
  sleep 2
done

echo "MySQL 已就绪，开始检查并创建数据库..."

# Step 2: 初始化数据库
# 给与可执行的权限
chmod +x ./script/docker/initdb/wechat-db.sh

./script/docker/initdb/wechat-db.sh

# Step 3: 启动 wechat 服务
docker-compose -f docker-compose.yml up -d
