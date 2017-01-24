#!/bin/bash

if [ "$1" == "cross" ]; then
    export PATH="/opt/data/x86_64-linux/usr/bin/arm-poky-linux-gnueabi:$PATH"
    . /opt/poky/2.2/environment-setup-cortexa7hf-neon-vfpv4-poky-linux-gnueabi
    CROSSCOMPILE=arm-poky-linux-gnueabi-
fi

PWD=`pwd`
GOPATH=$GOPATH:$PWD

FILES="github.com/astaxie/beego github.com/beego/bee "
FILES+=" code.google.com/p/graphics-go/graphics"
FILES+=" goconfig/config golanger.com/log golanger.com/utils"
FILES+=" jk/jkcommon jk/jkconfig jk/jkimage jk/jklog jk/jkmath jk/jkprotocol jk/jkparsedoc"
FILES+=" labix.org/v2/mgo labix.org/v2/mgo/bson"
FILES+=" github.com/tyranron/daemonigo/"
#FILES+=" github.com/jeffallen/mqtt github.com/surgemq/surgemq github.com/surgemq/surgemq/service"
#FILES+=" github.com/deckarep/gosx-notifier"
FILES+=" jk/jkeasycrypto"

# go get golang.org/x/mobile/cmd/gomobile
# gomobile init
# go get -d golang.org/x/mobile/example/basic
# gomobile build -target=android golang.org/x/mobile/example/basic
# gomobile install golang.org/x/mobile/example/basic
# gomobile build -target=ios golang.org/x/mobile/example/basic
# ios-deploy -b basic.app
# go get -d golang.org/x/mobile/example/bind/...
# git clone https://github.com/beego/i18n.git

for i in $FILES
do
    echo "${CROSSCOMPILE}go install $i ..."
    ${CROSSCOMPILE}go install $i
done

echo "Install Success !"
