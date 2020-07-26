#!/bin/sh
git clone https://github.com/charlesren/eagle.git
cd eagle/cmd/eagle-quote
go build quote.go
mv quote build/
cd build
docker build -t quote .
rm quote