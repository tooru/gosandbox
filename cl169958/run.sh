#!/bin/sh

TAG=cl169958

docker build . -t $TAG

docker run $TAG go version
docker run $TAG go run /root/main.go

echo

docker run $TAG /root/go/bin/go version
docker run $TAG /root/go/bin/go run /root/main.go
