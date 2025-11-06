package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/marie-mattheo/crm-project/internal/storage"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mini CRM v3!")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add a contact")
		fmt.Println("2. List contacts")
		fmt.Println("3. Update a contact")
		fmt.Println("4. Delete a contact")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		choice := readUserChoice(reader)

		switch choice {
		case 3:
			handleUpdateContact(reader, store)
		case 4:
			handleDeleteContact(reader, store)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// Les fonctions "handle..." s'occupent de l'interaction avec l'utilisateur
// et appellent la couche de stockage (le store) pour effectuer les opérations.
// Elles sont découplées du stockage : elles fonctionnent avec N'IMPORTE quel Storer.

func HandleAddContact(name string, email string) {
	// Le reste de la logique d'ajout
	contact := &storage.Contact{Name: name, Email: email}
	err = store.Add(contact)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(" Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to update: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName, err := readLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail, err := readLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}

func handleDeleteContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the contact to delete: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Contact with ID %d has been deleted.\n", id)
}

func readLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		// Affiche erreur de lecture (EOF, etc.)
		fmt.Printf("Error during read: %v\n", err)
		return "", err
	}

	trimmedInput := strings.TrimSpace(input)
	return trimmedInput, nil
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := readLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrChoice, err := strconv.Atoi(choice)
	if err != nil {
		return -1
	}
	return nbrChoice
}

func readInteger(reader *bufio.Reader) int {
	id, err := readLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrId, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	return nbrId
}
