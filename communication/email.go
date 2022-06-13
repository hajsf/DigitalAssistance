package communication

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
)

func SendEmail(sender, pushname, subject string) {
	user := "Kottouf.Procurement@googlemail.com" // "O365 logging name"
	password := "kvuyhakkojwjwvlx"               //"O365 logging pasword"
	// Generate App password
	// https://myaccount.google.com/u/3/security?rapt=AEjHL4Mi_H31usxNz7LtUye3Ao4XEAdxHYf-YTPJQdqh7bNwllrbfbNKnQT1f3P7Zo9nyLXQHkEff6TG7gnoFOjFUaIf92DvbQ
	// https://myaccount.google.com/u/3/apppasswords?rapt=AEjHL4OMZmrC7jubo9TOKRKEN3nqhgZNXqbkyudhBaum7vl4pqs8jsrev2pGQSjDfhiW3_omqrvSjJ9QIgcLFEgsNgDs2GvEuw
	smtpHost := "smtp.gmail.com" //"smtp.office365.com" // mail.kottouf.sa
	smtpPort := 587              //587 // 465 993 incoming

	d := mail.NewDialer(smtpHost, smtpPort, user, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := mail.NewMessage()
	m.SetHeader("From", "Kottouf.Procurement@googlemail.com") //
	m.SetHeader("To", "hasan.y@kottouf.net")
	m.SetHeader("Subject", fmt.Sprintf("%s: from %s - %s", subject, pushname, sender))
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("failed to send email:", err)
		// panic(err)
	}
	fmt.Println("Email sent")
}
