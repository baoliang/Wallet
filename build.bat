@echo off
:: build on windows
:: set GO111MODULE=ON
rd /s/q bin
mkdir bin
go build -o bin/wallet.exe
go build -o bin/wallet-cli cmd/masswalletcli/main.go