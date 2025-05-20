package models

import (
	"log"
	"sync"
)

type Client struct {
	ID    int
	Name  string
	Email string
}

type ClientStore interface {
	Create(client Client) Client
	All() []Client
	Get(id int) (Client, bool)
	Update(id int, client Client) bool
	Delete(id int) bool
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

func (s *InMemoryClientStore) Create(client Client) Client {
	log.Println("Server xxxxxxxxx")
	s.mu.Lock()
	defer s.mu.Unlock()
	client.ID = s.nextID
	log.Println(client)

	s.data[s.nextID] = client
	s.nextID++
	return client
}

func (s *InMemoryClientStore) All() []Client {
	log.Println("Server ddddddddd")

	s.mu.Lock()
	defer s.mu.Unlock()
	clients := []Client{}
	for _, c := range s.data {
		clients = append(clients, c)
	}
	return clients
}

func (s *InMemoryClientStore) Get(id int) (Client, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	c, ok := s.data[id]
	return c, ok
}

func (s *InMemoryClientStore) Update(id int, client Client) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[id]; !exists {
		return false
	}
	client.ID = id
	s.data[id] = client
	return true
}

func (s *InMemoryClientStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[id]; !exists {
		return false
	}
	delete(s.data, id)
	return true
}
