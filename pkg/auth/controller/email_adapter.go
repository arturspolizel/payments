package controller

type EmailAdapter struct {
}

func NewEmailAdapter() *EmailAdapter {
	return &EmailAdapter{}
}

func (a *EmailAdapter) SendEmail(address, content string) (err error) {
	// Mock email sending
	panic("Not implemented")
}
