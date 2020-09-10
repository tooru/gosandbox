#!/bin/sh

PRJ_ROOT=`dirname $0`/..

PATHSPEC=$1

if [ z$PATHSPEC = z ]; then
    PATHSPEC=master
fi

cd $PRJ_ROOT

docker build . --build-arg PATHSPEC=$PATHSPEC --build-arg='GOSSAFUNC=ssa'

