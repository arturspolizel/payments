package controller

import "github.com/rs/zerolog/log"

type EmailAdapter struct {
}

func NewEmailAdapter() *EmailAdapter {
	return &EmailAdapter{}
}

func (a *EmailAdapter) SendEmail(address, content string) (err error) {
	// Mock email sending
	log.Info().Msgf("Sending email to address %s, content: %s", address, content)
	return nil
}
