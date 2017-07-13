#!/bin/bash

type=$1
cross=${2-"amd64"}
basepath=`pwd`

function log() {
    echo "$*"
}

VERSION=`date +%Y%m%d%H%M%S`
PKG_NAME=$type-$cross-$VERSION.tar.bz2

case $cross in
    amd64)
    OS=linux
    ARCH=amd64
    bin_path=$basepath/bin/simpleserver
    CROSS_ENV="GOOS=linux GOARCH=amd64"
    ;;
    corei7|386)
    OS=linux
    ARCH=386
    bin_path=$basepath/bin/simpleserver
    CROSS_ENV="GOOS=linux GOARCH=386"
    ;;
    windows)
    OS=windows
    ARCH=amd64
    bin_path=$basepath/bin/simpleserver
    CROSS_ENV="GOOS=linux GOARCH=amd64"
    ;;
esac

case $type in
    simpleserver)
    rm -rf $bin_path
    if [ ! -f $bin_path ]; then
        log "build GOOS=$OS GOARCH=$ARCH"
        cd $basepath/src/jkprog/simpleserver
        GOOS=$OS GOARCH=$ARCH go build simpleserver.go
        cd -
        mv $basepath/src/jkprog/simpleserver/simpleserver $bin_path
    fi
    if [ ! -f $bin_path ]; then
        log "Error: simpleserver not build, check it"
        exit 2
    fi
    tmp_path=/tmp/simpleserver
    mkdir -p $tmp_path $tmp_path/bin
    cp -r $basepath/html $tmp_path
    cp -r $bin_path $tmp_path/bin
    cd /tmp
    tar cjvf $PKG_NAME simpleserver
    cd -
    mv /tmp/$PKG_NAME .
    ;;
esac