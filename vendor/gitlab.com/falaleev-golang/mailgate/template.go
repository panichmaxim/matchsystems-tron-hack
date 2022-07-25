package mailgate

// EmailTemplate mail template
type EmailTemplate struct {
	subject  string
	bodyHTML string
	bodyTXT  string
	to       []string
}

// Subject email subject
func (e *EmailTemplate) Subject() string {
	return e.subject
}

// BodyHTML email html body
func (e *EmailTemplate) BodyHTML() string {
	return e.bodyHTML
}

// BodyTXT email txt body
func (e *EmailTemplate) BodyTXT() string {
	return e.bodyTXT
}

// To receiver email address
func (e *EmailTemplate) To() []string {
	return e.to
}

// Template base email template
type Template interface {
	Subject() string
	BodyHTML() string
	BodyTXT() string
	To() []string
}
