#!/bin/bash

echo "Jun Chat Deploy...."
echo "Current Path:"
#输出本地路径
pwd
echo "Building Gateway..."
cd gateway
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gateway main.go
echo "Building Connect..."
cd ../connect
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o connect main.go
echo "Building Core..."
cd ../core
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o core main.go
echo "Building Queue..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o queue  main.go
#返回上层目录
cd ..

mkdir "JunChatServer"

echo "Deploy Sucess"
rm -rf "JunChatServer"
