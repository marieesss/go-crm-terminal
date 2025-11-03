package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	type Contact struct {
		Name    string
		Surname string
		Phone   string
		Email   string
	}

	contacts := []Contact{
		{Name: "Alice", Surname: "Smith", Phone: "123-456-7890", Email: "alice@example.com"},
		{Name: "Bob", Surname: "Johnson", Phone: "987-654-3210", Email: "bob@example.com"},
		{Name: "Charlie", Surname: "Brown", Phone: "555-555-5555", Email: "charlie@example.com"},
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
			fmt.Print("Nom de l'utilisateur : ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Println("Utilisateur ajouté :", name)
		case "2":
			fmt.Println("Liste des utilisateurs")
		case "3":
			fmt.Println("Utilisateur supprimé")
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
