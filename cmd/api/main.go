package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
	"github.com/marieesss/go-crm-terminal/internal/handler"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var contacts []domain.Contact

	contacts = append(contacts, domain.Contact{Name: "Alice", Surname: "Smith", Phone: "123-456-7890", Email: "alice@example.com"})
	contacts = append(contacts, domain.Contact{Name: "Bob", Surname: "Johnson", Phone: "987-654-3210", Email: "bob@example.com"})
	contacts = append(contacts, domain.Contact{Name: "Charlie", Surname: "Brown", Phone: "555-555-5555", Email: "charlie@example.com"})

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Modifier un contact")
		fmt.Println("5. Quitter l'appli")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			contact, err := handler.AddUser(&contacts)
			if err != nil {
				fmt.Println("Erreur lors de l'ajout de l'utilisateur :", err)
			} else {
				fmt.Println("Utilisateur ajouté :", contact)
			}
		case "2":
			handler.ListUsers(&contacts)
		case "3":
			removed, err := handler.DeleteUser(&contacts)
			if err != nil {
				fmt.Println("Erreur lors de la suppression de l'utilisateur :", err)
			} else {
				fmt.Println("Utilisateur supprimé :", removed.Name, removed.Surname)
			}
		case "4":
			fmt.Println("Utilisateur modifié")
		case "5":
			fmt.Println("Fermeture de l'application...")
			return
		default:
			fmt.Println("Option inconnue ❌")
		}
		fmt.Println()
	}
}
