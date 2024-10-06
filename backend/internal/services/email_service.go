package services

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"
)

type EmailService struct {
	dialer *mail.Dialer
}

func NewEmailService() *EmailService {
	d := mail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	return &EmailService{dialer: d}
}

func (s *EmailService) SendInvitation(to, gatheringName, inviteLink string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("FROM_EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("You're invited to %s!", gatheringName))

	body := fmt.Sprintf(`
		<html>
			<body>
				<h2>You're invited to %s!</h2>
				<p>You've been invited to join a gathering. Click the link below to RSVP:</p>
				<a href="%s">RSVP Now</a>
				<p>We hope to see you there!</p>
			</body>
		</html>
	`, gatheringName, inviteLink)

	m.SetBody("text/html", body)

	return s.dialer.DialAndSend(m)
}
