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
cd ../queue
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o queue  main.go
#返回上层目录
cd ..

mkdir "JunChatServer"

mv gateway/gateway ./JunChatServer
mv connect/connect ./JunChatServer
mv core/core ./JunChatServer
mv queue/queue ./JunChatServer
cp ./config.yaml ./JunChatServer

echo "Compress..."
tar -czf ./JunChatServer.tar.gz ./JunChatServer/
echo "Compress Success"
echo "Upload File...."
scp ./JunChatServer.tar.gz root@182.92.239.63:/root
echo "Upload File Success"

rm -rf "JunChatServer"
rm -rf "JunChatServer.tar.gz"


###云端
ssh root@182.92.239.63 < ./deploy_cloud.sh
