package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

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
			fmt.Println("Utilisateur ajouté")
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
