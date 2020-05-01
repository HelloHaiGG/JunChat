#!/bin/bash

cd ~/root
rm -rf JunChatServer/
echo "Decompression..."
tar -xzvf JunChatServer.tar.gz
echo "Decompression Success"

echo "stop core..."
ps -ef | grep ./core | grep -v grep | awk '{print $2}' | xargs kill -9
echo "stop connect..."
ps -ef | grep ./connect | grep -v grep | awk '{print $2}' | xargs kill -9
echo "stop queue..."
ps -ef | grep ./queue | grep -v grep | awk '{print $2}' | xargs kill -9
echo "stop gateway..."
ps -ef | grep ./gateway | grep -v grep | awk '{print $2}' | xargs kill -9

echo "start gateway..."
cd JunChatServer
./gateway 2>&1 >/root/logs/gateway.log &
echo "start connect..."
./connect 2>&1 >/root/logs/connect &
echo "start core..."
./core 2>&1 >/root/logs/core.log &
echo "start queue..."
./queue 2>&1 >/root/logs/queue &

echo "Deploy Success"

exit