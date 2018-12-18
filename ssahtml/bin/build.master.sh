#!/bin/sh

docker build . --build-arg PATHSPEC=master --build-arg='GOSSAFUNC=ssa:*+'
