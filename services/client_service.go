package services

import (
	"errors"
	"go-client-app/models"
	"go-client-app/validators"
	"net/mail"
	"strconv"
)

type ClientService struct {
	store models.ClientStore
}

func NewClientService(store models.ClientStore) *ClientService {
	return &ClientService{store}
}

func (s *ClientService) CreateClient(name, email string) error {
	if err := validators.ValidateClient(name, email); err != nil {
		return err
	}

	// Optional: parse to ensure valid email format
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New("invalid email format")
	}

	return s.store.Create(models.Client{
		Name:  name,
		Email: email,
	})
}

func (s *ClientService) DeleteClientByID(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return errors.New("invalid client ID")
	}
	return s.store.Delete(id)
}

func (s *ClientService) ListClients() []models.Client {
	return s.store.All()
}
