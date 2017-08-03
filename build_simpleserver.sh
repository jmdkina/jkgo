#!/bin/bash

file=$1

[ -z $file ] && {
    log "Usage: $0 file to build"
    log "  Unsupport install libs"
    exit 1
}

function log() {
    echo $*
}

go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GOVERSION=`go version`'" $file

log "move ${file#*/} bin"
mv ${file#*/} bin

exit 0
