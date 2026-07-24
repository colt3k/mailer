package message

import (
	"strings"

	"github.com/colt3k/mail"
)

// BuildMessage constructs a mail.Message from the provided parameters.
func BuildMessage(from, to, cc, ccname, subject, msg string, html bool, filePath string) *mail.Message {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	if len(cc) > 0 && len(ccname) > 0 {
		m.SetAddressHeader("Cc", cc, ccname)
	} else if len(cc) > 0 {
		m.SetAddressHeader("Cc", cc, cc)
	}
	m.SetHeader("Subject", subject)
	if html {
		m.SetBody("text/html", msg)
	} else {
		m.SetBody("text/plain", msg)
	}
	if len(strings.TrimSpace(filePath)) > 0 {
		m.Attach(filePath)
	}
	return m
}
