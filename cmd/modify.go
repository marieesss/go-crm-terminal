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
	modifyID    int
	modifyName  string
	modifyEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Met à jour un contact.",
	Long:  `La commande 'update' permet de modifier les informations d'un contact existant.`,
	Run: func(cmd *cobra.Command, args []string) {
		if modifyName == "" || modifyEmail == "" || modifyID == 0 {
			handleUpdateContact(store)
			return
		}

		handleUpdatedContact(modifyName, modifyEmail, modifyID, store)
	},
}

func handleUpdateContact(store storage.Storer) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ID of the contact to update: ")
	modifyID := utils.ReadInteger(reader)
	if modifyID == -1 {
		return
	}

	// On vérifie que le contact existe avant de demander les nouvelles infos
	existingContact, err := store.GetByID(modifyID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingContact.Name)

	fmt.Printf("New name (%s): ", existingContact.Name)
	newName, err := utils.ReadLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("New email (%s): ", existingContact.Email)
	newEmail, err := utils.ReadLine(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	handleUpdatedContact(newName, newEmail, existingContact.ID, store)

}

func handleUpdatedContact(name string, email string, id int, store storage.Storer) {
	err := store.Update(id, name, email)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully.")
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Définition des drapeaux pour la commande 'update'
	updateCmd.Flags().StringVarP(&modifyName, "name", "n", "", "Nom du contact")
	updateCmd.Flags().StringVarP(&modifyEmail, "email", "e", "", "Email du contact")

}
