#!/bin/bash

goget="github.com/astaxie/beego github.com/beego/bee github.com/Unknwon/goconfig"

for i in $goget
do
   do="go get $i"
   echo "$do"
   go get $i
done
