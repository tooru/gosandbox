#!/bin/sh

docker build . --build-arg PATHSPEC=go1.11.4 --build-arg='GOSSAFUNC=ssa'

