package main

import (
	"github.com/marie-mattheo/crm-project/cmd"
	"github.com/marie-mattheo/crm-project/internal/storage"
)

func main() {
	var store = storage.NewMemoryStore()
	cmd.Execute(store)
}
