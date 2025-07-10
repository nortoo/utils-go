package ses

type Content struct {
	// Subject is the subject of an email.
	Subject string

	// HTML represents the HTML format content of the email.
	HTML string

	// Text represents the plain text format content of the email.
	Text string
}
