#!/bin/sh

TAG=cl169958

docker build . -t $TAG --build-arg PATHSPEC=4145c5da1f533fafd928769d18d5be60968cb9dc

docker run $TAG go version
docker run $TAG go run /root/main.go

echo

docker run $TAG /root/go/bin/go version
docker run $TAG /root/go/bin/go run /root/main.go
