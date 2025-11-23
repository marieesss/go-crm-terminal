package main

import (
	"fmt"

	"github.com/marie-mattheo/crm-project/cmd"
	"github.com/marie-mattheo/crm-project/internal/storage"
)

func main() {
	store, err := storage.NewJSONStore("contacts.json")
	if err != nil {
		fmt.Printf("Erreur storage: %v\n", err)
	}
	cmd.Execute(store)
}
