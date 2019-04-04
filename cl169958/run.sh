#!/bin/sh

docker build . -t cl169958

docker run cl169958 go run /root/main.go
docker run cl169958 /root/go/bin/go run /root/main.go
