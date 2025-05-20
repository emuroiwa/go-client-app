package tests

import (
	"strconv"
	"testing"

	"go-client-app/models"
	"go-client-app/services"
)

func setupTestService() *services.ClientService {
	store := models.NewInMemoryClientStore()
	return services.NewClientService(store)
}

func TestCreateClient_ValidData(t *testing.T) {
	service := setupTestService()

	err := service.CreateClient("Test User", "test@example.com")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	clients := service.ListClients()
	if len(clients) != 1 {
		t.Errorf("expected 1 client, got %d", len(clients))
	}

	if clients[0].Name != "Test User" || clients[0].Email != "test@example.com" {
		t.Error("client data does not match")
	}
}

func TestCreateClient_InvalidEmail(t *testing.T) {
	service := setupTestService()

	err := service.CreateClient("Invalid Email User", "not-an-email")
	if err == nil {
		t.Error("expected error for invalid email, got nil")
	}
}

func TestDeleteClient_ValidID(t *testing.T) {
	service := setupTestService()

	_ = service.CreateClient("User", "user@email.com")

	clients := service.ListClients()
	if len(clients) == 0 {
		t.Fatal("no clients created")
	}
	id := strconv.Itoa(clients[0].ID)
	err := service.DeleteClientByID(id)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(service.ListClients()) != 0 {
		t.Error("expected client to be deleted")
	}
}

func TestDeleteClient_InvalidID(t *testing.T) {
	service := setupTestService()

	err := service.DeleteClientByID("999")
	if err == nil {
		t.Error("expected error for invalid ID, got nil")
	}
}
