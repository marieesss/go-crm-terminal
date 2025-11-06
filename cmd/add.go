package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marie-mattheo/crm-project/internal/storage"
	"github.com/marie-mattheo/crm-project/internal/utils"
	"github.com/spf13/cobra"
)

var (
	name  string
	email string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un nouveau contact à un fichier JSON de configuration.",
	Long: `La commande 'add' permet d'ajouter un nouveau contact (nom, email)
à un fichier JSON spécifié. Si le fichier n'existe pas, il sera créé.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" || email == "" {
			handleAddContact(store)
			return
		}

		handleStoreContacts(name, email, store)
	},
}

func handleAddContact(store storage.Storer) {
	var (
		reader *bufio.Reader
		name   string
		email  string
		err    error
	)

	reader = bufio.NewReader(os.Stdin)

	// Boucle pour le nom (obligatoire)
	for {
		fmt.Print("Enter contact name: ")
		name, err = utils.ReadLine(reader)
		if err != nil {
			// Erreur système lors de la lecture (EOF, etc.)
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if name != "" {
			break // L'entrée est valide (non vide)
		}
		fmt.Println("Name cannot be empty. Please enter a value.")
	}

	// Boucle pour l'email (obligatoire)
	for {
		fmt.Print("Enter contact email: ")
		email, err = utils.ReadLine(reader)
		if err != nil {
			fmt.Printf("Read error: %v. Please try again.\n", err)
			continue
		}
		if email != "" {
			break // L'entrée est valide (non vide)
		}
		fmt.Println("Email cannot be empty. Please enter a value.")
	}

	handleStoreContacts(name, email, store)

}

func handleStoreContacts(name string, email string, store storage.Storer) {
	contact := &storage.Contact{Name: name, Email: email}
	err := store.Add(contact)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf(" Contact '%s' added with ID %d.\n", contact.Name, contact.ID)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Définition des drapeaux pour la commande 'add'
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom du contact")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email du contact")

}
