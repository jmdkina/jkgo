@echo off
set PWD=%cd%
set GOPATH=%GOPATH%;%PWD%

set installfiles=(code.google.com/p/graphics-go/graphics github.com/astaxie/beego github.com/beego/bee ^
	github.com/go-sql-driver/mysql github.com/howeyc/fsnotify github.com/jtolds/gls ^
	goconfig/config golanger.com/log golanger.com/utils ^
	jk/jkcommon jk/jkconfig jk/jkimage jk/jklog jk/jkmath jk/jknetwork jk/jkparsedoc jk/jkprotocol jk/jkserver ^
	labix.org/v2/mgo labix.org/v2/mgo/bson ^
	helper bveth  github.com/jeffallen/mqtt github.com/surgemq/surgemq github.com/surgemq/surgemq/service ^
	jkprog/jkhttpserver jkprog/jkencoderimg jkprog/jktransfer jkprog/jkcli jkprog/jkmng)

for %%i in %installfiles% do (
	echo install %%i
	go install %%i
)

pause
