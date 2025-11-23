package storage

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GORMStore stocke les contacts dans une base SQLite via GORM
type GORMStore struct {
	db *gorm.DB
}

// NewGORMStore crée un nouveau GORMStore et initialise la base de données
func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migration : GORM crée automatiquement la table "contacts" à partir de la struct Contact
	err = db.AutoMigrate(&Contact{})
	if err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

// Add ajoute un nouveau contact dans la base de données
func (g *GORMStore) Add(contact *Contact) error {
	result := g.db.Create(contact)
	if result.Error != nil {
		return result.Error
	}
	// L'ID est automatiquement rempli par GORM après l'insertion
	return nil
}

// GetAll récupère tous les contacts de la base de données
func (g *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	result := g.db.Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}
	return contacts, nil
}

// GetByID récupère un contact par son ID
func (g *GORMStore) GetByID(id int) (*Contact, error) {
	var contact Contact
	result := g.db.First(&contact, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrContactNotFound(id)
		}
		return nil, result.Error
	}
	return &contact, nil
}

// Update modifie un contact existant
func (g *GORMStore) Update(id int, newName string, newEmail string) error {
	contact, err := g.GetByID(id)
	if err != nil {
		return err
	}

	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}

	result := g.db.Save(contact)
	return result.Error
}

// Delete supprime un contact par son ID
func (g *GORMStore) Delete(id int) error {
	result := g.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrContactNotFound(id)
	}
	return nil
}
