#!/bin/bash

cd ~/web
rm -rf "dist"
echo "Decompression..."
tar -xzvf dist.tar.gz
echo "Decompression Success"
rm -rf "dist.tar.gz"
echo "Success!"