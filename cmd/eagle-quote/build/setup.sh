#!/bin/sh
git clone https://github.com/charlesren/eagle.git
cd eagle/cmd/eagle-quote
go build quote.go
docker build -t quote .