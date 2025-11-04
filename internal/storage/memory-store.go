package storage

import (
	"github.com/marieesss/go-crm-terminal/internal/domain"
)

type MemoryStore struct {
	data   map[int]*domain.Contact
	nextID int
}

var contacts = map[int]*domain.Contact{
	1: {Name: "Alice", Surname: "Smith", Phone: "123", Email: "alice@example.com"},
	2: {Name: "Bob", Surname: "Jones", Phone: "456", Email: "bob@example.com"},
	3: {Name: "Charlie", Surname: "Brown", Phone: "789", Email: "charlie@example.com"},
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data:   contacts,
		nextID: len(contacts) + 1,
	}
}

func (m *MemoryStore) Add(c *domain.Contact) error {
	m.data[m.nextID] = c
	m.nextID++
	return nil
}

func (m *MemoryStore) GetAll() []*domain.Contact {
	list := []*domain.Contact{}
	for _, c := range m.data {
		list = append(list, c)
	}
	return list

}
