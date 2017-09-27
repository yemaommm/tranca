cd ..
set APPNAME=windowmain.go
set GOPATH=%GOPATH%;%cd%
go run run.go go run %APPNAME% -host="0.0.0.0:3000"