#!/bin/sh
git clone https://github.com/charlesren/eagle.git
cd eagle/cmd/quote/config
GOOS=linux go build ../quote.go
docker build -t quote:v1 .
rm -f quote