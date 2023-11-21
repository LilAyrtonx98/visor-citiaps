#!/bin/bash
echo "STARTING BENCHMARK TESTS"
go test ./tests/ -cpuprofile cpu.prof -memprofile mem.prof -bench . -run=^B -benchtime 1s -count 1
go tool pprof -png cpu.prof
go tool pprof -png mem.prof
rm ./cpu.prof
rm ./mem.prof
echo "ENDING BENCHMARK TESTS"