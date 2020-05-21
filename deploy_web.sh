#!/bin/bash

echo "Jun Chat Web Deploy..."
echo "local path:"
pwd
echo "Building..."
npm run build:prod
echo "Building Success!"

echo "Compress..."
tar -czf ./dist.tar.gz ./dist/
echo "Compress Success"
echo "Upload File...."
scp ./dist.tar.gz root@182.92.239.63:/root/web
echo "Upload File Success"

rm -rf "dist"
rm -rf "dist.tar.gz"

###云端
ssh root@182.92.239.63 < ./deploy_cloud_web.sh
# root123456.