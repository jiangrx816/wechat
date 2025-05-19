#!/bin/bash

echo "等待 MySQL 启动完成..."

# 等待 mysql 容器的健康检查通过
until docker inspect --format "{{json .State.Health.Status }}" mysql-db | grep -q '"healthy"'; do
  sleep 2
done

echo "MySQL 已就绪，开始检查并创建数据库..."

# 执行数据库初始化（注意此处是直接通过 docker exec 调用）
docker exec -i mysql-db mysql -uroot -p123456 <<EOF
CREATE DATABASE IF NOT EXISTS wechat DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
EOF

echo "数据库创建完成"
