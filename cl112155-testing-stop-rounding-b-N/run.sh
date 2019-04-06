#!/bin/sh

TAG=cl169958

docker build . -t $TAG --build-arg PATHSPEC=03a79e94ac72f2425a6da0b399dc2f660cf295b6

docker run $TAG go version
docker run $TAG go test -bench=. /root/foo_test.go

echo

docker run $TAG /root/go/bin/go version
docker run $TAG /root/go/bin/go test -bench=. /root/foo_test.go
