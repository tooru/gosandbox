#!/bin/sh

docker build . --build-arg PATHSPEC=go1.9.2 --build-arg='GOSSAFUNC=ssa'
