@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end


:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w src

go get code.google.com/p/goprotobuf/proto
go get code.google.com/p/goprotobuf/protoc-gen-go
go get github.com/go-sql-driver/mysql
go get github.com/codegangsta/martini
go get github.com/golang/groupcache
go get github.com/goraft/raft
go get code.google.com/p/log4go

:end
echo finished
