@echo off

for %%i in (protocol.proto) do c:\protoc.exe --go_out=. %%i
pause