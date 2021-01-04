package main

import "fmt"

// Example - Builder Parameter

type email struct {
	from, to, subject, body string
}

type action func(builder *EmailBuilder)

//EmailBuilder struct
type EmailBuilder struct {
	email email
}

func sendEmail(email *email) {
	fmt.Println("email sent ...")
	fmt.Printf("To:%s,\nSubject:%s\n%s\nRegards,\n%s\n", email.to, email.subject, email.body, email.from)
}

//SendEmail function is for client to send email
func SendEmail(action action) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmail(&builder.email)
}

//To sets the email's "To" address
func (eb *EmailBuilder) To(value string) *EmailBuilder {
	eb.email.to = value
	return eb
}

//From sets the email's "From" address
func (eb *EmailBuilder) From(value string) *EmailBuilder {
	eb.email.from = value
	return eb
}

//Body sets the email's "Body"
func (eb *EmailBuilder) Body(value string) *EmailBuilder {
	eb.email.body = value
	return eb
}

//Subject sets the email's "Subject"
func (eb *EmailBuilder) Subject(value string) *EmailBuilder {
	eb.email.subject = value
	return eb
}

//RunExample2 to run example on builder parameters
func RunExample2() {
	SendEmail(func(b *EmailBuilder) {
		b.
			To("World@gmail.com").
			From("me@gmail.com").
			Subject("Hello world").
			Body("Sample body for a dummy email")
	})
}
