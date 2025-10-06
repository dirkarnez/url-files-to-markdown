@echo off

set DOWNLOADS_DIR=%USERPROFILE%\Downloads

set GOROOT=%DOWNLOADS_DIR%\go1.21.0.windows-amd64\go
set GOPATH=%DOWNLOADS_DIR%\gopath

set PATH=^
%WINDIR%\System32;^
%GOROOT%\bin;

SET PATH=^
%PATH%;^
%DOWNLOADS_DIR%\PortableGit\bin;^
%DOWNLOADS_DIR%\winlibs-x86_64-posix-seh-gcc-11.2.0-mingw-w64-9.0.0-r1\mingw64;^
%DOWNLOADS_DIR%\winlibs-x86_64-posix-seh-gcc-11.2.0-mingw-w64-9.0.0-r1\mingw64\bin;

go build main.go &&^
pause

