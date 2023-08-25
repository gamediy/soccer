#!/bin/bash
read -p "请输入版本号: " ver
name="soccer-admin"
image_name="$name:$ver"
# 强制拉群 最近 GIT 代码
git fetch --all
git reset --hard origin/main

# 编译 Go 代码
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin main.go
# 构建 Docker 镜像
docker build -t "$image_name" .

# 删除旧的 Docker 容器
docker rm -f "$name" >/dev/null 2>&1
# 运行 Docker 容器
docker run -d --name "$name" -p 9000:9000 "$image_name" --gf.gcfg.file="config-dev.yaml"
