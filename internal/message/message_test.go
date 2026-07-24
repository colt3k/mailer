package message

import (
	"bytes"
	"strings"
	"testing"
)

func TestBuildMessagePlainText(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "", "", "subj", "body", false, "")

	from := m.GetHeader("From")
	if len(from) != 1 || from[0] != "a@b.com" {
		t.Errorf("From = %v, want [a@b.com]", from)
	}

	to := m.GetHeader("To")
	if len(to) != 1 || to[0] != "c@d.com" {
		t.Errorf("To = %v, want [c@d.com]", to)
	}

	subj := m.GetHeader("Subject")
	if len(subj) != 1 || subj[0] != "subj" {
		t.Errorf("Subject = %v, want [subj]", subj)
	}

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if !strings.Contains(out, "text/plain") {
		t.Errorf("expected text/plain body, got:\n%s", out)
	}
	if !strings.Contains(out, "body") {
		t.Errorf("expected body content, got:\n%s", out)
	}
}

func TestBuildMessageHTML(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "", "", "subj", "<b>body</b>", true, "")

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if !strings.Contains(out, "text/html") {
		t.Errorf("expected text/html body, got:\n%s", out)
	}
	if !strings.Contains(out, "<b>body</b>") {
		t.Errorf("expected html body content, got:\n%s", out)
	}
}

func TestBuildMessageCCWithName(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "cc@e.com", "CcName", "subj", "body", false, "")

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if !strings.Contains(out, "cc@e.com") {
		t.Errorf("expected cc address in output, got:\n%s", out)
	}
	if !strings.Contains(out, "CcName") {
		t.Errorf("expected cc name in output, got:\n%s", out)
	}
}

func TestBuildMessageCCNoName(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "cc@e.com", "", "subj", "body", false, "")

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if !strings.Contains(out, "cc@e.com") {
		t.Errorf("expected cc address in output, got:\n%s", out)
	}
}

func TestBuildMessageNoCC(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "", "", "subj", "body", false, "")

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if strings.Contains(out, "Cc:") {
		t.Errorf("expected no Cc header, got:\n%s", out)
	}
}

func TestBuildMessageWhitespaceFilePath(t *testing.T) {
	m := BuildMessage("a@b.com", "c@d.com", "", "", "subj", "body", false, "   ")

	var buf bytes.Buffer
	m.WriteTo(&buf)
	out := buf.String()
	if strings.Contains(out, "application/octet-stream") {
		t.Errorf("expected no attachment for whitespace path, got:\n%s", out)
	}
}
