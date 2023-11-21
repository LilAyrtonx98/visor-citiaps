package tests

import (
	"os"
	"testing"

	"github.com/citiaps/visor-mapas-backend/auth"
	"github.com/citiaps/visor-mapas-backend/db"
	"github.com/citiaps/visor-mapas-backend/utils"
	"github.com/gin-gonic/gin"
)

// Router de GIN para testing de routes_*_test.go
var router_routes *gin.Engine

// Router de GIN para testing de auth_*_test.go
var router_auth *gin.Engine

func TestMain(m *testing.M) {
	// Preparar requisitos
	utils.LoadConfig("../config/config.test.yml")
	utils.InitLogger()
	db.Setup()
	utils.InitSMTPServer()
	auth.InitAuthSession()

	// Definir un router de GIN para routes_*_test.go
	router_routes = gin.Default()

	// Definir un router de GIN para routes_*_test.go
	router_auth = gin.Default()

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

// Benchmark test

var result int

func Factorial(n int) int {
	var result int
	if n == 0 || n == 1 {
		return 1
	}
	result = Factorial(n-1) * n
	return result
}

func benchmarkFact(b *testing.B, num int) {
	var y int
	for x := 0; x < b.N; x++ {
		//run my awesome test method
		y = Factorial(num)
		//fmt.Printf("Y = %d\n", y) -- uncomment out to get the first results in "take one" below
	}
	result = y
}

func BenchmarkFact20(b *testing.B) {
	benchmarkFact(b, 200)
}
