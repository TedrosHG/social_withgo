package mailer

import (
	"bytes"
	"errors"
	"learn-go/internal/env"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("API key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m *mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	sandboxUsername := env.GetString("SANDBOX_USERNAME", "")
	sandboxPassword := env.GetString("SANDBOX_PASSWORD", "")
	// template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject.String())
	message.AddAlternative("text/html", body.String())

	// Set up the SMTP dialer
	// dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)

	// For Sandbox, the host is usually sandbox.smtp.mailtrap.io
	// and the username is NOT "api" (it's a unique ID found in your inbox settings)
	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, sandboxUsername, sandboxPassword)

	if err := dialer.DialAndSend(message); err != nil {
		return -1, err
	}

	return 200, nil
}
