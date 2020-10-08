@echo off
:: windows gen pb file
del /F/A/Q  *.pb.go
protoc -I=. -I=%GOPATH%\src -I=%GOPATH%\src\github.com\gogo\protobuf\protobuf  --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go  --gogo_out=.\  config.proto
