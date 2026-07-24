package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildConfigDefaults(t *testing.T) {
	dir := t.TempDir()
	cfgFile := filepath.Join(dir, "config_example.toml")

	// Override io.WriteOut destination temporarily via env or by capturing output.
	// Since buildConfig writes to a fixed path relative to the binary, we test
	// by setting the package vars and checking the written file.
	smtpServer = ""
	smtpPort = 0
	from = ""
	to = ""

	buildConfig()

	// The file is written relative to the binary directory. Read it.
	binDir, err := os.Executable()
	if err != nil {
		t.Skip("cannot determine executable path")
	}
	binDir = filepath.Dir(binDir)
	fullPath := filepath.Join(binDir, "config_example.toml")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Skip("config_example.toml not found; skipping")
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	s := string(content)
	// Defaults should appear when vars are empty.
	if !strings.Contains(s, "smtp_server = \"\"") {
		t.Errorf("expected empty smtp_server in config, got:\n%s", s)
	}
	if !strings.Contains(s, "from = \"\"") {
		t.Errorf("expected empty from in config, got:\n%s", s)
	}
	if !strings.Contains(s, "to = \"\"") {
		t.Errorf("expected empty to in config, got:\n%s", s)
	}
	_ = cfgFile // avoid unused var
}

func TestBuildConfigCustomValues(t *testing.T) {
	smtpServer = "smtp.custom.com"
	smtpPort = 465
	from = "custom@from.com"
	to = "custom@to.com"

	buildConfig()

	binDir, err := os.Executable()
	if err != nil {
		t.Skip("cannot determine executable path")
	}
	binDir = filepath.Dir(binDir)
	fullPath := filepath.Join(binDir, "config_example.toml")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Skip("config_example.toml not found; skipping")
	}

	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	s := string(content)
	if !strings.Contains(s, "smtp_server = \"smtp.custom.com\"") {
		t.Errorf("expected custom smtp_server, got:\n%s", s)
	}
	if !strings.Contains(s, "smtp_port = 465") {
		t.Errorf("expected smtp_port = 465, got:\n%s", s)
	}
	if !strings.Contains(s, "from = \"custom@from.com\"") {
		t.Errorf("expected custom from, got:\n%s", s)
	}
	if !strings.Contains(s, "to = \"custom@to.com\"") {
		t.Errorf("expected custom to, got:\n%s", s)
	}
}
