@echo off
chcp 65001
REM 设置交叉编译参数
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

REM 编译生成 Linux 可执行文件
@REM go build -o main

REM 设置交叉编译参数
set GOOS=windows
set GOARCH=amd64

REM 编译生成 Windows 可执行文件
go build -o kfgpt_serve3.exe

echo 编译完成！
pause
