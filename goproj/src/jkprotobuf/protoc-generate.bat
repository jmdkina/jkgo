@echo off

for %%i in (protocol.proto) do protoc --go_out=. %%i
pause