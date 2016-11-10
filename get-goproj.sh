#!/bin/bash

goget="github.com/astaxie/beego github.com/beego/bee github.com/Unknwon/goconfig github.com/rthornton128/goncurses"

for i in $goget
do
   do="go get $i"
   echo "$do"
   go get $i
done
