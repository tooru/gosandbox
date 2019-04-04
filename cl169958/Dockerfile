FROM golang:1.12.1-alpine3.9

COPY src/main.go /root/

RUN apk add --no-cache git bash gcc libc-dev

WORKDIR /root
RUN git clone https://github.com/golang/go.git

ARG PATHSPEC=4145c5da1f533fafd928769d18d5be60968cb9dc

ENV GOROOT_BOOTSTRAP=/usr/local/go

WORKDIR go/src
RUN git checkout $PATHSPEC
RUN ./make.bash
