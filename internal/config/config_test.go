package config

import (
	"errors"
	"testing"
)

func TestSMTPConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     SMTPConfig
		wantErr error
	}{
		{
			name: "valid config",
			cfg: SMTPConfig{
				Server: "smtp.example.com",
				Port:   587,
				From:   "sender@example.com",
				To:     "recipient@example.com",
			},
			wantErr: nil,
		},
		{
			name:    "missing server",
			cfg:     SMTPConfig{Port: 587, From: "a@b.com", To: "c@d.com"},
			wantErr: errors.New("smtp server required"),
		},
		{
			name:    "missing from",
			cfg:     SMTPConfig{Server: "smtp.example.com", Port: 587, To: "c@d.com"},
			wantErr: errors.New("from required"),
		},
		{
			name:    "missing to",
			cfg:     SMTPConfig{Server: "smtp.example.com", Port: 587, From: "a@b.com"},
			wantErr: errors.New("to required"),
		},
		{
			name:    "port too low",
			cfg:     SMTPConfig{Server: "smtp.example.com", Port: 0, From: "a@b.com", To: "c@d.com"},
			wantErr: errors.New("port must be between 1 and 65535"),
		},
		{
			name:    "port too high",
			cfg:     SMTPConfig{Server: "smtp.example.com", Port: 65536, From: "a@b.com", To: "c@d.com"},
			wantErr: errors.New("port must be between 1 and 65535"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cfg.Validate()
			if (got == nil) != (tt.wantErr == nil) {
				t.Errorf("Validate() error = %v, want %v", got, tt.wantErr)
			}
			if got != nil && tt.wantErr != nil && got.Error() != tt.wantErr.Error() {
				t.Errorf("Validate() error = %q, want %q", got.Error(), tt.wantErr.Error())
			}
		})
	}
}
