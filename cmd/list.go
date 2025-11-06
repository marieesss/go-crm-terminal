package cmd

import (
	"fmt"

	"github.com/marie-mattheo/crm-project/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste tous les contacts.",
	Long:  `La commande 'list' permet d'afficher tous les contacts présents.`,
	Run: func(cmd *cobra.Command, args []string) {
		handleListContacts(store)
	},
}

func handleListContacts(store storage.Storer) {
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println(" No contacts to display.")
		return
	}

	fmt.Println("\n--- Contact List ---")
	for _, contact := range contacts {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
