#!/bin/bash

echo "创建 go.mod 和 go.sum 文件"
go mod init myproject

echo "创建目录结构"
mkdir -pv cmd/ internal/api/ internal/pkg/ pkg/mypubliclib/
touch cmd/main.go                  \
      internal/api/server.go       \
      internal/api/server_test.go  \
      internal/pkg/util.go         \
      pkg/mypubliclib/mypubliclib.go
# 提示初始化完成
echo "项目初始化完成"
