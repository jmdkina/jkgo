@echo off
set PWD=%cd%
set GOPATH=%GOPATH%;%PWD%

for %%i in (golanger.com\config golanger.com\framework\web golanger.com\i18n golanger.com\log golanger.com\middleware golanger.com\utils golanger.com\session\filesession golanger.com\session\cookiesession golanger.com\session\memorysession golanger.com\database\activerecord goconfig\config code.google.com\p\go.net\websocket code.google.com\p\goprotobuf\proto code.google.com\p\graphics-go\graphics code.google.com\p\goprotobuf\protoc-gen-go jk\jkcommon jk\jklog jk\jkprotobuf jk\jkserver  helper bveth labix.org\v2\mgo labix.org\v2\mgo\bson) do go install %%i
pause
