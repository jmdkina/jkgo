#!/bin/bash

goget="github.com/astaxie/beego github.com/beego/bee"

for i in $goget
do
   do="go get $i"
   echo "$do"
   go get $i
done
