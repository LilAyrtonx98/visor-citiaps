#!/bin/bash
echo "STARTING UNIT TESTS"
go test -v -coverpkg ./... ./tests/ -coverprofile=coverage.out -run=^T
go tool cover -html=coverage.out -o coverage.html
echo "ENDING UNIT TESTS"
echo "COVERAGE REPORT: coverage.html"