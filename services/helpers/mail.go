package services;

import (
    "log"
    "net/smtp"

    configs "github.com/markdingemanse/loveless/configs"
)

func Communicate(body string) {
    from := configs.Config("LOVELESS_MAIL_USER")
    pass := configs.Config("LOVELESS_MAIL_PASSWORD")
    to := configs.Config("LOVELESS_MAIL_DESTINATION")
    subject := configs.Config("LOVELESS_MAIL_SUBJECT")
    server := configs.Config("LOVELESS_MAIL_URL")
    port := configs.Config("LOVELESS_MAIL_PORT")

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject +"\n\n" +
        body

    err := smtp.SendMail(server + ":" + port,
        smtp.PlainAuth("", from, pass, server),
        from, []string{to}, []byte(msg))

    if err != nil {
        log.Printf("smtp error: %s", err)
        return
    }
}
