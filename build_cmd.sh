#! /bin/bash

mkdir -p ./main
touch ./main/main.go

echo "package main" > ./main/main.go
echo "import \"github.com/Alonza0314/logger-go/cmd\"" >> ./main/main.go
echo "func main() {" >> ./main/main.go
echo "	cmd.Execute()" >> ./main/main.go
echo "}" >> ./main/main.go

GOOS=linux GOARCH=amd64 go build -o ./build/logger-go-linux-amd64 ./main/main.go
GOOS=windows GOARCH=amd64 go build -o ./build/logger-go-windows-amd64.exe ./main/main.go
GOOS=darwin GOARCH=amd64 go build -o ./build/logger-go-darwin-amd64 ./main/main.go
GOOS=darwin GOARCH=arm64 go build -o ./build/logger-go-darwin-arm64 ./main/main.go

rm -rf ./main