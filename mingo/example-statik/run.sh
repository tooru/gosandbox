#!/bin/sh

go generate || exit 1
go build || exit 1
