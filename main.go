package main

import (
	"fmt"

	"github.com/marie-mattheo/crm-project/cmd"
	"github.com/marie-mattheo/crm-project/internal/storage"
)

func main() {
	store, err := storage.NewGORMStore("contacts.db")
	if err != nil {
		fmt.Printf("Erreur storage: %v\n", err)
		return
	}
	cmd.Execute(store)
}
