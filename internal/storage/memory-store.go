package storage

import (
	"fmt"

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

func (m *MemoryStore) GetAll() map[int]*domain.Contact {
	return m.data
}

func (m *MemoryStore) Delete(id int) error {
	if _, ok := m.data[id]; !ok {
		return fmt.Errorf("aucun contact trouvé avec l'id %d", id)
	}
	delete(m.data, id)
	return nil
}

func (m *MemoryStore) Update(id int, c *domain.Contact) error {
	if _, ok := m.data[id]; !ok {
		return fmt.Errorf("aucun contact trouvé avec l'id %d", id)
	}
	// on remplace l'entrée (simple et clair)
	m.data[id] = c
	return nil
}
