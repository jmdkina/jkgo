#!/bin/bash 
goget="github.com/astaxie/beego github.com/beego/bee github.com/Unknwon/goconfig github.com/rthornton128/goncurses"
goget+=" github.com/gogap/logs"
goget+=" github.com/henrylee2cn/faygo github.com/henrylee2cn/fay github.com/gorilla/websocket"
goget+=" gopkg.in/mgo.v2"

for i in $goget
do
   if [ -d src/$i ]; then
       continue
   fi
   do="go get $i"
   echo "$do"
   go get $i
done

CUR=`pwd`

srcs="github.com/alecthomas/log4go.git"
srcs+=" github.com/golang/protobuf.git"
srcs+=" github.com/gorilla/websocket.git"
srcs+="  github.com/matishsiao/goInfo"
srcs+=" github.com/beego/i18n"
srcs+=" github.com/kardianos/service"

for src in $srcs
do
    log4godir=${src%.*}
    if [ ! -d src/$log4godir ]; then
        pre=${log4godir%/*} 
        mkdir -p src/$pre
        cd src/$pre
        echo "git clone https://$src"
        git clone https://$src
        cd $CUR
        echo "go install $log4godir"
        go install $log4godir
    fi
done
