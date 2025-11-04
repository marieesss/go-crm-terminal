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

	contacts := map[int]*domain.Contact{
		1: {Name: "Alice", Surname: "Smith", Phone: "123", Email: "alice@example.com"},
		2: {Name: "Bob", Surname: "Jones", Phone: "456", Email: "bob@example.com"},
		3: {Name: "Charlie", Surname: "Brown", Phone: "789", Email: "charlie@example.com"},
	}

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
			if err := handler.ModifyUser(&contacts); err != nil {
				fmt.Printf("Erreur lors de la modification : %v\n", err)
			}
		case "5":
			fmt.Println("Fermeture de l'application...")
			return
		default:
			fmt.Println("Option inconnue ❌")
		}
		fmt.Println()
	}
}
