#!/bin/bash

goget="github.com/astaxie/beego github.com/beego/bee github.com/Unknwon/goconfig github.com/rthornton128/goncurses"

for i in $goget
do
   do="go get $i"
   echo "$do"
   go get $i
done

git clone https://github.com/golang/protobuf.git src/github.com/golang/protobuf
git clone https://github.com/gorilla/websocket.git src/github.com/gorilla/websocket
