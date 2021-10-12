#!/bin/sh

echo "====================="
echo "format source files"
echo "go fmt ./..."
go fmt ./...
echo "====================="
echo "run go vet ./..."
echo "====================="
go vet ./...
echo "====================="
echo "start basic lint"
echo "====================="
golint ./...
echo "====================="
echo "check source files"
echo "====================="
golangci-lint run
