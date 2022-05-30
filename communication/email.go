package communication

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
)

func SendEmail(sender, pushname, subject string) {
	user := "Kottouf.Procurement@googlemail.com" // "O365 logging name"
	password := "Dana0Yara"                      //"O365 logging pasword"
	smtpHost := "smtp.gmail.com"                 //"smtp.office365.com" // mail.kottouf.sa
	smtpPort := 587                              //587 // 465 993 incoming

	d := mail.NewDialer(smtpHost, smtpPort, user, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := mail.NewMessage()
	m.SetHeader("From", "Kottouf.Procurement@googlemail.com") //
	m.SetHeader("To", "hasan.y@kottouf.net")
	m.SetHeader("Subject", fmt.Sprintf("%s: from %s - %s", subject, pushname, sender))
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
