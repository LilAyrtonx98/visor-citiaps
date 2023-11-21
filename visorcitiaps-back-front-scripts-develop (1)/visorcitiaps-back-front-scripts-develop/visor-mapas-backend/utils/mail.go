package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type SmtpServer struct {
	Host string
	Port string
}

func (server *SmtpServer) ServerName() string {
	return server.Host + ":" + server.Port
}

var SmtpServerConfig SmtpServer

func InitSMTPServer() {
	SmtpServerConfig = SmtpServer{Host: Config.Email.Smtp, Port: Config.Email.Port}
}

func NewMail(sender string, to []string, subject string, body string) *Mail {
	mail := new(Mail)
	mail.Sender = sender
	mail.To = to
	mail.Subject = subject
	mail.Body = body
	return mail
}

func (mail *Mail) Build() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}

func (mail *Mail) Send() bool {

	messageBody := mail.Build()

	auth := smtp.PlainAuth("", mail.Sender, Config.Email.Password, SmtpServerConfig.Host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SmtpServerConfig.Host,
	}

	conn, err := tls.Dial("tcp", SmtpServerConfig.ServerName(), tlsconfig)
	if err != nil {
		log.Println(err)
		return false
	}

	client, err := smtp.NewClient(conn, SmtpServerConfig.Host)
	if err != nil {
		log.Println(err)
		return false
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Println(err)
		return false
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Println(err)
		return false
	}
	for _, k := range mail.To {
		if err = client.Rcpt(k); err != nil {
			log.Println(err)
			return false
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Println(err)
		return false
	}

	err = w.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	client.Quit()

	log.Println("Mail enviado exitosamente a:", mail.To)
	return true
}

//utils.CreateMail([]string{"destinatario@gmail.com"}, "Notificación", "Esta es una prueba")
func CreateMail(to []string, subject string, body string) bool {
	var value bool
	go func() {
		value = NewMail(Config.Email.Sender, to, subject, body).Send()
	}()
	return value
}
