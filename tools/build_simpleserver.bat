
@echo off
set file=simpleserver/simpleserver

set DATETIME=%date%-%time%
set GOVERSION=1.11
set GOPATH=%GOPATH%;%cd%

@echo on
go build -ldflags "-X main.VERSION=1.0.0 -X 'main.BUILD_TIME=%DATETIME%' -X 'main.GOVERSION=%GOVERSION%'" %file%

move simpleserver.exe bin
