package infrastructure

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KarnerTh/query-lookout/config"
	"github.com/KarnerTh/query-lookout/usecase/notify"
)

//go:embed mail_template.html
var mailTemplateContent string

type mailData struct {
	notify.Notification
	DeepLink string
}

type mailNotifier struct {
	fromAddress  string
	toAddress    string
	smtpHost     string
	smtpPort     string
	baseUrl      string
	mailTemplate *template.Template
}

func NewMailNotifier(config config.Config) notify.Notifier {
	tmpl, err := template.New("mail_template").Parse(mailTemplateContent)
	if err != nil {
		log.WithError(err).Fatal("Could not load mail template")
	}

	return mailNotifier{
		fromAddress:  config.MailFromAddress(),
		toAddress:    config.MailToAddress(),
		smtpHost:     config.MailSmtpHost(),
		smtpPort:     config.MailSmtpPort(),
		baseUrl:      config.BaseUrl(),
		mailTemplate: tmpl,
	}
}

func (n mailNotifier) Send(value notify.Notification) error {
	if err := n.validateMailConfig(); err != nil {
		return err
	}

	smtpAddress := fmt.Sprintf("%s:%s", n.smtpHost, n.smtpPort)
	to := strings.Split(n.toAddress, ",")
	msg := n.getMailMsg(value)

	if err := smtp.SendMail(smtpAddress, nil, n.fromAddress, to, msg); err != nil {
		log.WithError(err).Error("Could not send mail")
		return err
	}

	return nil
}

func (n mailNotifier) getMailMsg(value notify.Notification) []byte {
	fromHeader := fmt.Sprintf("From: %s\n", n.fromAddress)
	toHeader := fmt.Sprintf("To: %s\n", n.toAddress)
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := getMailContent(n.mailTemplate, value, n.baseUrl)

	return []byte(toHeader + fromHeader + subject + mime + body)
}

func getMailContent(template *template.Template, data notify.Notification, baseUrl string) string {
	inputData := mailData{
		Notification: data,
		DeepLink:     fmt.Sprintf("%s/lookout/%d", baseUrl, data.LookoutId),
	}

	bodyBuf := new(bytes.Buffer)
	if err := template.Execute(bodyBuf, inputData); err != nil {
		log.WithError(err).Error("Could not load mail template")
		return ""
	}

	return bodyBuf.String()
}

func (n mailNotifier) validateMailConfig() error {
	missingConfigs := []string{}
	if n.fromAddress == "" {
		missingConfigs = append(missingConfigs, "mail_from_address")
	}

	if n.toAddress == "" {
		missingConfigs = append(missingConfigs, "mail_to_address")
	}

	if n.smtpHost == "" {
		missingConfigs = append(missingConfigs, "mail_smtp_host")
	}

	if n.smtpPort == "" {
		missingConfigs = append(missingConfigs, "mail_smtp_port")
	}

	if len(missingConfigs) > 0 {
		return fmt.Errorf("Config missing: %s", strings.Join(missingConfigs, ", "))
	}

	return nil
}
