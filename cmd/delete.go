package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marie-mattheo/crm-project/internal/storage"
	"github.com/marie-mattheo/crm-project/internal/utils"
	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprime un contact.",
	Long:  `La commande 'delete' permet de supprimer un contact existant par son ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == 0 {
			handleDeleteContact(store)
			return
		}

		err := store.Delete(deleteID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Contact with ID %d has been deleted.\n", deleteID)
	},
}

func handleDeleteContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the ID of the contact to delete: ")
	id := utils.ReadInteger(reader)
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

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Flag pour fournir l'ID directement
	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID du contact à supprimer")
}
