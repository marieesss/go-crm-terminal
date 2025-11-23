package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// JSONStore est une implémentation de Storer
// qui persiste les contacts dans un fichier JSON.
type JSONStore struct {
	filePath string
	contacts map[int]*Contact
	nextID   int
}

func NewJSONStore(filePath string) (*JSONStore, error) {
	js := &JSONStore{
		filePath: filePath,
		contacts: make(map[int]*Contact),
		nextID:   1,
	}

	if err := js.loadFromFile(); err != nil {
		// Si le fichier n'existe pas, on continue avec un store vide
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("erreur lors du chargement du fichier JSON: %w", err)
		}
	}

	return js, nil
}

type jsonContact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (js *JSONStore) loadFromFile() error {
	data, err := os.ReadFile(js.filePath)
	if err != nil {
		return err
	}

	var list []jsonContact
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, c := range list {
		contact := &Contact{
			ID:    c.ID,
			Name:  c.Name,
			Email: c.Email,
		}
		js.contacts[c.ID] = contact
		if c.ID >= js.nextID {
			js.nextID = c.ID + 1
		}
	}

	return nil
}

func (js *JSONStore) saveToFile() error {
	list := make([]jsonContact, 0, len(js.contacts))
	for _, c := range js.contacts {
		list = append(list, jsonContact{
			ID:    c.ID,
			Name:  c.Name,
			Email: c.Email,
		})
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(js.filePath, data, 0644)
}

func (js *JSONStore) Add(contact *Contact) error {
	contact.ID = js.nextID
	js.contacts[contact.ID] = contact
	js.nextID++
	return js.saveToFile()
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	var all []*Contact
	for _, c := range js.contacts {
		all = append(all, c)
	}
	return all, nil
}

func (js *JSONStore) GetByID(id int) (*Contact, error) {
	contact, ok := js.contacts[id]
	if !ok {
		return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}
	return contact, nil
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	contact, err := js.GetByID(id)
	if err != nil {
		return err
	}

	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}

	return js.saveToFile()
}

func (js *JSONStore) Delete(id int) error {
	if _, ok := js.contacts[id]; !ok {
		return fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}
	delete(js.contacts, id)
	return js.saveToFile()
}
