Rem mv to some where
@echo off

set mvfile=%1
set distdir=C:\proj\jkgo\goproj\bin\

echo "mv %mvfile% %distdir%"
move %mvfile% %distdir%
