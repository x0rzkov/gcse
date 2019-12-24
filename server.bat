go install github.com/x0rzkov/gcse/server
@if errorlevel 1 goto exit
%GOPATH%\bin\server

:exit
