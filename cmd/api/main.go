package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/handler"
	"github.com/marieesss/go-crm-terminal/internal/storage"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	memoryStore := storage.NewMemoryStore()

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1.⭐ Ajouter un contact")
		fmt.Println("2.⭐ Lister les contacts")
		fmt.Println("3.⭐ Supprimer un contact")
		fmt.Println("4.⭐ Modifier un contact")
		fmt.Println("5.⭐ Quitter l'appli")
		fmt.Print("Choisissez une option : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			contact, err := handler.AddUser(memoryStore, reader)
			if err != nil {
				fmt.Println("Erreur lors de l'ajout de l'utilisateur :", err)
			} else {
				fmt.Println("😍 Utilisateur ajouté :", contact)
			}
		case "2":
			handler.ListUsers(memoryStore)
		case "3":
			removed, err := handler.DeleteUser(memoryStore, reader)
			if err != nil {
				fmt.Println("Erreur lors de la suppression de l'utilisateur :", err)
			} else {
				fmt.Println("Utilisateur supprimé :", removed.Name, removed.Surname)
			}
		case "4":
			updated, id, err := handler.UpdateUser(memoryStore, reader)
			if err != nil {
				fmt.Println("Erreur:", err)
				break
			}
			fmt.Printf("🎉 Contact %d mis à jour: %s %s\n", id, updated.Name, updated.Surname)
		case "5":
			fmt.Println("Fermeture de l'application...")
			return
		default:
			fmt.Println("Option inconnue ❌")
		}
		fmt.Println()
	}
}
