package email

import (
	"net/smtp"
	"os"
)

// SMTPServer data to configure server connection
type SMTPServer struct {
	Host string
	Port string
}

// ServerName URI to smtp server
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// BuildMessage creates the message to be sent
func BuildMessage(from, to, subject, body string) []byte {
	headers := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n"
	return []byte(headers + body)
}

// SendEmail sends an email
func SendEmail(from, password, to, subject, body string) error {
	// Retrieve SMTP server info from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	smtpServer := SMTPServer{Host: smtpHost, Port: smtpPort}

	message := BuildMessage(from, to, subject, body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.Host)

	// Sending email.
	err := smtp.SendMail(smtpServer.ServerName(), auth, from, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}
