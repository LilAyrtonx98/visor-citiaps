#!/bin/bash
echo "STARTING GOLANG BUILD"
go build -o visor-mapas-backend app.go
echo "ENDING GOLAND BUILD"
echo "BINARY FILE: visor-mapas-backend"