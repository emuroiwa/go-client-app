package models

import (
	"fmt"
	"sync"
)

type Client struct {
	ID    int
	Name  string
	Email string
}

type ClientStore interface {
	Create(Client) error
	Delete(id int) error
	All() []Client
}
type InMemoryClientStore struct {
	data   map[int]Client
	nextID int
	mu     sync.Mutex
}

func NewInMemoryClientStore() *InMemoryClientStore {
	return &InMemoryClientStore{
		data:   make(map[int]Client),
		nextID: 1,
	}
}

func (s *InMemoryClientStore) All() []Client {
	s.mu.Lock()
	defer s.mu.Unlock()
	clients := []Client{}
	for _, c := range s.data {
		clients = append(clients, c)
	}
	return clients
}

func (s *InMemoryClientStore) Create(client Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	client.ID = s.nextID
	s.data[s.nextID] = client
	s.nextID++
	return nil
}

func (s *InMemoryClientStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[id]; !exists {
		return fmt.Errorf("client with ID %d not found", id)
	}
	delete(s.data, id)
	return nil
}
