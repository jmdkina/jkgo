#!/bin/bash

goget="github.com/astaxie/beego github.com/beego/bee github.com/Unknwon/goconfig github.com/rthornton128/goncurses"
goget+=" github.com/gogap/logs"

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
srcs+="github.com/golang/protobuf.git"
srcs+="github.com/gorilla/websocket.git"

for src in $srcs
do
    dir=${log4go%.*}
    if [ ! -d src/$log4godir ]; then
        pre=${log4go%/*} 
        mkdir -p src/$pre
        cd src/$pre
        git clone https://$log4go
        cd $CUR
    fi
done
