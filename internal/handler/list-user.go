package handler

import (
	"fmt"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func ListUsers(memoryStore domain.Storer) {
	fmt.Println("\n=== Liste des contacts ===")
	contacts := memoryStore.GetAll()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	for id, contact := range contacts {
		fmt.Printf("\nContact %d:\n", id)
		fmt.Printf("Nom: %s\n", contact.Name)
		fmt.Printf("Prénom: %s\n", contact.Surname)
		fmt.Printf("Téléphone: %s\n", contact.Phone)
		fmt.Printf("Email: %s\n", contact.Email)
		fmt.Println("------------------------")
	}
}
