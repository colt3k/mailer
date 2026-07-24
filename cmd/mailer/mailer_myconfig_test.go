package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/colt3k/mycli"
)

func TestLoadMyconfigToml(t *testing.T) {
	// Walk up from cmd/mailer/ to the project root where pkgr/ lives.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get working directory: %v", err)
	}
	root := filepath.Dir(filepath.Dir(wd))
	configPath := filepath.Join(root, "pkgr", "myconfig.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Skipf("config file not found: %s", configPath)
	}

	tw := mycli.Toml()
	if err := tw.LoadToml(configPath); err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	tests := []struct {
		key   string
		value string
	}{
		{"smtp_server", "smtp.gmail.com"},
		{"smtp_port", "587"},
		{"from", "gcstang@gmail.com"},
		{"smtp_username", "gcstang@gmail.com"},
		{"to", "gcstang@gmail.com"},
		{"tls", "true"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if !tw.Has(tt.key) {
				t.Fatalf("expected key %q in config", tt.key)
			}
			got := tw.Get(tt.key)
			if got == nil {
				t.Fatalf("got nil for key %q", tt.key)
			}
			var gotStr string
			switch v := got.(type) {
			case string:
				gotStr = v
			case int64:
				gotStr = strconv.FormatInt(v, 10)
			case bool:
				gotStr = strconv.FormatBool(v)
			default:
				gotStr = strings.TrimSpace(fmt.Sprint(v))
			}
			if gotStr != tt.value {
				t.Errorf("key %q = %q, want %q", tt.key, gotStr, tt.value)
			}
		})
	}
}
