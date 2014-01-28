@echo off

protoc.exe  --go_out = . *.proto

echo finished
