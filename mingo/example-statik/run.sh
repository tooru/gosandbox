#!/bin/sh

statik || exit 1
go build || exit 1
