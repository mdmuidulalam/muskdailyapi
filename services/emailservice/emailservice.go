package emailService

import (
	"net/smtp"
	"strings"

	config "muskdaily.com/config"
)

func SendMail(to []string, from string, subject string, body string) {
	configuration := config.GetConfiguration()

	if from == "" {
		from = configuration.Smtp.UserName
	}

	message := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", from, configuration.Smtp.Password, configuration.Smtp.Host)
	err := smtp.SendMail(configuration.Smtp.Host+":"+configuration.Smtp.Port, auth, from, to, message)

	if err != nil {
		panic(err)
	}
}
