#!/bin/bash

type=$1
cross=${2-"amd64"}
basepath=`pwd`

function log() {
    echo "$*"
}

VERSION=`date +%Y%m%d%H%M%S`
PKG_NAME=$type-$VERSION.tar.bz2

case $cross in
    amd64)
    OS=linux
    ARCH=amd64
    bin_path=$basepath/bin/simpleserver
    ;;
    corei7|386)
    OS=linux
    ARCH=386
    bin_path=$basepath/linux_386/bin/simpleserver
    ;;
    windows)
    OS=windows
    ARCH=amd64
    bin_path=$basepath/bin/simpleserver
    ;;
esac

case $type in
    simpleserver)
    if [ ! -f $bin_path ]; then
        ${CROSS_ENV} go install simpleserver
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