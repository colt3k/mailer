package config

import "errors"

// SMTPConfig holds the SMTP connection and sender/recipient parameters.
type SMTPConfig struct {
	Server string
	Port   int64
	User   string
	Pass   string
	From   string
	To     string
}

// Validate returns an error if required fields are missing or port is out of range.
func (c SMTPConfig) Validate() error {
	if c.Server == "" {
		return errors.New("smtp server required")
	}
	if c.From == "" {
		return errors.New("from required")
	}
	if c.To == "" {
		return errors.New("to required")
	}
	if c.Port < 1 || c.Port > 65535 {
		return errors.New("port must be between 1 and 65535")
	}
	return nil
}
