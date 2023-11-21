package tests

import (
	"testing"

	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestHassPassword(t *testing.T) {
	// Nombre test
	test := "HashPassword"

	// Ejecutar test
	_, err := utils.HashPassword("password")

	//Verificar resultado
	failed := err != nil

	//Escribir resultado
	utils.Test(t, failed, test)
}

func TestCheckPasswordHash(t *testing.T) {
	test := "CheckPasswordHash"

	passwordHash, err := utils.HashPassword("password")

	failed := err != nil || !utils.CheckPasswordHash("password", passwordHash)

	utils.Test(t, failed, test)
}
