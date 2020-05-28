#!/bin/bash

cd ~
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

echo "del redis connected record..."
/root/mid/redis-5.0.7/src/redis-cli -h 182.92.239.63 -p 6379 del LIVE:ON:SERVER

echo "start gateway..."
cd JunChatServer
./gateway 2>&1 >/root/logs/gateway.log &
echo "start connect..."
#./connect 2>&1 >/root/logs/connect-1.log &
./connect -NET node-1 -RPC node-1 2>&1 >/root/logs/connect-1.log &
./connect -NET node-2 -RPC node-2 2>&1 >/root/logs/connect-2.log &
./connect -NET node-3 -RPC node-3 2>&1 >/root/logs/connect-3.log &
./connect -NET node-4 -RPC node-4 2>&1 >/root/logs/connect-4.log &
echo "start core..."
./core 2>&1 >/root/logs/core.log &
echo "start queue..."
./queue 2>&1 >/root/logs/queue.log &

echo "Deploy Success"

exit