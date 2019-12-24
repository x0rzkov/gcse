go install github.com/x0rzkov/gcse/indexer
@if errorlevel 1 goto exit
%GOPATH%\bin\indexer

:exit
