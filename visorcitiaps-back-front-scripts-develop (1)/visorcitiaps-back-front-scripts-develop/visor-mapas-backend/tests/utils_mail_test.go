package tests

import (
	"testing"

	"github.com/citiaps/visor-mapas-backend/utils"
)

func TestInitSMTPServer(t *testing.T) {
	test := "InitSMTPServer"

	failed := utils.SmtpServerConfig.Host == "" || utils.SmtpServerConfig.Port == ""

	utils.Test(t, failed, test)
}

// func TestCreateMail(t *testing.T) {
// 	test := "CreateMail"

// 	mail := "cesar.kreep@gmail.com"

// 	failed := !utils.NewMail(utils.Config.Email.Sender, []string{mail}, "Notificación de prueba", "Esta es una prueba").Send()
// 	utils.Test(t, failed, test)
// }
