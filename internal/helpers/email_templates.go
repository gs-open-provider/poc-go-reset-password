package helpers

import "github.com/sendgrid/sendgrid-go/helpers/mail"

// RPFromAddress Constant
var RPFromAddress = mail.NewEmail("GSOP", "noreply@gsop.com")

// RPEmailSubject Constant
const RPEmailSubject = "[GSOP] Please Reset your Password"

// RPEmailContent Constant
const RPEmailContent = "Please follow the instructions to reset your password"

// GetResetPasswordTemplate Function
func GetResetPasswordTemplate(link string) string {
	return `
		We heard that you lost your GSOP password. Sorry about that!
		<br /><br />
		But don’t worry! You can use the following link to reset your password:
		<br /><br />
		` + link + `
		<br /><br />
		If you don’t use this link within 30 mins, it will expire. To get a new password reset link, visit http://localhost:9090/index
		<br /><br />
		Thanks,
		<br />
		Your friends at GSOP
	`
}
