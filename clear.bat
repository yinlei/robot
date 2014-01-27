@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end


:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

./bin/protoc-gen-go.exe 

:end
echo finished
