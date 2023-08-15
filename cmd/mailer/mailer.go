package main

import (
	"bytes"
	"github.com/colt3k/mailer/internal/update"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/colt3k/mail"
	"github.com/colt3k/mycli"
	log "github.com/colt3k/nglog/ng"
	"github.com/colt3k/utils/config"
	"github.com/colt3k/utils/file"
	"github.com/colt3k/utils/io"
	"github.com/colt3k/utils/lock"
	"github.com/colt3k/utils/mathut"
	"github.com/colt3k/utils/version"
)

const (
	appName     = "mailer"
	title       = "Mailer"
	description = "Send mail via SMTP"
	author      = "Gil Collins"
	copyright   = "(c) 2021 ColTech Inc.,"
	companyName = "colt3k"
)

var (
	proxy      string
	smtpServer string
	smtpPort   int64
	smtpUser   string
	smtpPass   string
	from       string
	to         string
	cc         string
	ccname     string
	subject    = "Test Message"
	msg        = "Hello, test message"
	html       bool
	filePath   string

	skipUpdateCheck bool
	logDir          = file.HomeFolder()
	logFile         = filepath.Join(logDir, companyName, appName+".log")
	c               *mycli.CLI
	l               *lock.Lock
	locked          bool
)

var cfg = config.NewConfig()

func init() {
	l = lock.New(appName)
	if l.Try() {
		locked = true
	}
	// Init Lock file here!!!
	fa, err := log.NewFileAppender("*", logFile, "", 0)
	if err != nil {
		log.Logf(log.FATAL, "issue creating file appender\n%+v", err)
	}
	ca := log.NewConsoleAppender("*")
	log.Modify(log.LogLevel(log.INFO), log.ColorsOn(), log.Appenders(ca, fa))
}
func setLogger() error {
	if logDir != file.HomeFolder() || mycli.Debug {
		logFile = filepath.Join(logDir, companyName, appName+".log")
		// update our logger
		fa, err := log.NewFileAppender("*", logFile, "", 0)
		if err != nil {
			return err
		}
		ca := log.NewConsoleAppender("*")
		if mycli.Debug {
			log.Modify(log.LogLevel(log.DEBUG), log.ColorsOn(), log.Appenders(ca, fa))
		} else {
			log.Modify(log.LogLevel(log.INFO), log.ColorsOn(), log.Appenders(ca, fa))
		}
	}
	return nil
}
func setupFlags() {
	c = mycli.NewCli(nil, nil)
	c.Title = title
	c.Description = description
	c.Version = version.VERSION
	c.BuildDate = version.BUILDDATE
	c.GitCommit = version.GITCOMMIT
	c.GoVersion = version.GOVERSION
	c.Author = author
	c.Copyright = copyright
	c.PostGlblAction = func() error { return setLogger() }
	c.Flgs = []mycli.CLIFlag{
		&mycli.BoolFlg{Variable: &skipUpdateCheck, Name: "skip_update", ShortName: "skip", Usage: "set flag to skip update check", Value: false},
		&mycli.StringFlg{Variable: &logDir, Name: "log_dir", ShortName: "ld", Usage: "set logging directory", Value: file.HomeFolder()},

		&mycli.StringFlg{Variable: &from, Name: "from", Usage: "Who is the email from `user@domain.com`"},
		&mycli.StringFlg{Variable: &to, Name: "to", Usage: "Whom to send the email to `user@domain.com`"},
		&mycli.StringFlg{Variable: &cc, Name: "cc", Usage: "Whom to cc on the email `user@domain.com`"},
		&mycli.StringFlg{Variable: &ccname, Name: "ccname", Usage: "Name to cc on the email `John`"},
		&mycli.StringFlg{Variable: &subject, Name: "subject", ShortName: "s", Usage: "Subject of the Email", Value: subject},
		&mycli.BoolFlg{Variable: &html, Name: "html", Usage: "Message body default plain set for html"},
		&mycli.StringFlg{Variable: &msg, Name: "message", ShortName: "m", Usage: "Message body to send in either quoted plaintext or html", Value: msg},
		&mycli.StringFlg{Variable: &filePath, Name: "file", ShortName: "f", Usage: "Attach a file `FILE`"},

		// Hidden
		&mycli.StringFlg{Variable: &smtpServer, Name: "smtp_server", Usage: "SMTP Server", Value: "localhost"},
		&mycli.Int64Flg{Variable: &smtpPort, Name: "smtp_port", Usage: "SMTP Port", Value: 587},
		&mycli.StringFlg{Variable: &smtpUser, Name: "smtp_username", Usage: "SMTP Username"},
		&mycli.StringFlg{Variable: &smtpPass, Name: "smtp_password", Usage: "SMTP Password"},
	}

	c.Cmds = []*mycli.CLICommand{
		{
			Name:   "update",
			Usage:  "check for updates",
			Action: func() { update.CheckUpdate(appName) },
		},
		{
			Name:      "buildconfig",
			ShortName: "bc",
			Usage:     "build generic configuration you can fill in",
			Action:    func() { buildConfig() },
		},
		{
			Name:   "send",
			Usage:  "send email",
			Action: func() { run() },
		},
	}
	//Executes validation or processes input
	err := c.Parse()
	if err != nil {
		if locked {
			l.Unlock()
			log.Logln(log.DEBUG, "unlocked")
		}
		log.Logf(log.FATAL, "error(s)\n%+v", err)
	}
}
func main() {
	setupFlags()
	if locked {
		l.Unlock()
		log.Logln(log.DEBUG, "unlocked")
	}
	os.Exit(0)
}

func buildConfig() {
	var byt bytes.Buffer

	byt.WriteString("smtp_server=\"")
	byt.WriteString(smtpServer)
	byt.WriteString("\"\n")
	byt.WriteString("smtp_port=")
	byt.WriteString(mathut.FmtInt(int(smtpPort)))
	byt.WriteString("\n")
	byt.WriteString("smtp_username=\"")
	byt.WriteString(smtpUser)
	byt.WriteString("\"\n")
	byt.WriteString("smtp_password=\"")
	byt.WriteString(smtpPass)
	byt.WriteString("\"\n")
	byt.WriteString("from=\"")
	if len(from) > 0 {
		byt.WriteString(from)
	} else {
		byt.WriteString("user@domain.com")
	}
	byt.WriteString("\"\n")
	byt.WriteString("to=\"")
	if len(to) > 0 {
		byt.WriteString(to)
	} else {
		byt.WriteString("user@domain.com")
	}

	byt.WriteString("\"\n")

	io.WriteOut(byt.Bytes(), "config_example.toml")
}

func run() {
	if len(smtpServer) == 0 {
		log.Logln(log.FATAL, "smtp server required")
	}
	if len(from) == 0 {
		log.Logln(log.FATAL, "from required")
	}
	if len(to) == 0 {
		log.Logln(log.FATAL, "to required")
	}
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	//CC
	if len(cc) > 0 && len(ccname) > 0 {
		m.SetAddressHeader("Cc", cc, ccname)
	} else if len(cc) > 0 && len(ccname) <= 0 {
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

	d := mail.NewDialer(smtpServer, int(smtpPort), smtpUser, smtpPass)

	if smtpPort != 25 {
		d.StartTLSPolicy = mail.MandatoryStartTLS
	} else {
		d.SSL = false
		d.StartTLSPolicy = mail.NoStartTLS
	}
	d.Timeout = 20 * time.Second
	log.Logln(log.DEBUG, "Sending")
	if err := d.DialAndSend(m); err != nil {
		log.Logf(log.FATAL, "error sending mail\n%+v", err)
	}
	log.Logln(log.DEBUG, "Finished")
}
