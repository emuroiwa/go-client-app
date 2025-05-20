package validators

import (
	"errors"
	"net/mail"
)

func ValidateClient(name, email string) error {
	if name == "" || email == "" {
		return errors.New("Name and email are required")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("Invalid email format")
	}
	return nil
}
