FROM golang:1.11.3-alpine3.8

RUN apk add --no-cache git bash gcc libc-dev graphviz

WORKDIR /root
RUN git clone https://github.com/golang/go.git
WORKDIR go/src

ARG PATHSPEC=go1.9.2

ENV GOROOT_BOOTSTRAP=/usr/local/go
RUN git checkout $PATHSPEC
RUN ./make.bash
ENV PATH=/root/go/bin:$PATH
COPY src/main.go /root/
RUN apk add --no-cache graphviz

ARG GOSSAFUNC=main
RUN GOSSAFUNC=$GOSSAFUNC go build /root/main.go
