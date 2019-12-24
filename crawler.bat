go install github.com/x0rzkov/gcse/crawler
@if errorlevel 1 goto exit
%GOPATH%\bin\crawler

:exit
