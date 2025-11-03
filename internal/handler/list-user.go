package handler

import (
	"fmt"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func ListUsers(contacts []domain.Contact) {
	fmt.Println("\n=== Liste des contacts ===")
	if len(contacts) == 0 {
		fmt.Println("Aucun contact enregistré.")
		return
	}

	for i, contact := range contacts {
		fmt.Printf("\nContact %d:\n", i+1)
		fmt.Printf("Nom: %s\n", contact.Name)
		fmt.Printf("Prénom: %s\n", contact.Surname)
		fmt.Printf("Téléphone: %s\n", contact.Phone)
		fmt.Printf("Email: %s\n", contact.Email)
		fmt.Println("------------------------")
	}
}
