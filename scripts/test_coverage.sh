#!/bin/bash
go test --coverprofile=coverage.out ./... -coverpkg ./...
cat coverage.out | grep -v ".pb.go" | grep -v "main.go" > coverage.out
go tool cover --html=coverage.out
