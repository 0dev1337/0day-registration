package mailer

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func SendEmail(recipientEmail, name string) {
	smtpHost := "smtp.mail.me.com"
	smtpPort := 587

	var from1 string
	if strings.Contains(recipientEmail, "stu.adamasuniversity.ac.in") {
		from1 = "zeronighter@icloud.com"
	} else {
		from1 = "zero@zerodev.me"
	}
	from := ""
	password := "" // app-specific password

	subject := fmt.Sprintf("📬 Missed this? Join Club 0day, %s!", name)
	body := fmt.Sprintf(`Hi %s 👋,

Welcome to *Club 0day*! We're excited to have you with us. 🚀

You're now part of a growing community of builders and innovators. Here's how you can stay connected:

🔗 Join the community:  
- WhatsApp: https://chat.whatsapp.com/DxVzQ3XbMdk6KMsmapmHN9  
- Discord: https://discord.gg/VB4FMbuukE

💬 Need help or have questions?  
We're here for you — don't hesitate to reach out.

Let’s learn, build, and grow together. 💡

Best regards,  
Team 0day`, name)

	body = fmt.Sprintf(`Hi %s 👋,

It looks like some of our last emails didn’t reach everyone — so we're resending the important links to join the Club 0day community:

🔗 WhatsApp: https://chat.whatsapp.com/DxVzQ3XbMdk6KMsmapmHN9  
🔗 Discord: https://discord.gg/VB4FMbuukE

See you there! 🚀

– Team 0day`, name)
	msg := strings.Join([]string{
		fmt.Sprintf("From: %s", from1),
		fmt.Sprintf("To: %s", recipientEmail),
		fmt.Sprintf("Subject: %s", subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=\"utf-8\"",
		"",
		body,
	}, "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
		auth,
		from,
		[]string{recipientEmail},
		[]byte(msg),
	)

	if err != nil {
		log.Printf("Error sending email to %s (%s): %v", name, recipientEmail, err)
		return
	}

	log.Printf("Email sent successfully to %s (%s)", name, recipientEmail)
}
