package main

import (
	"fmt"
	"os"

	"github.com/marie-mattheo/crm-project/cmd"
	"github.com/marie-mattheo/crm-project/internal/config"
	"github.com/marie-mattheo/crm-project/internal/storage"
)

func main() {
	// Charger la configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Erreur chargement config: %v\n", err)
		os.Exit(1)
	}

	// Créer le store approprié selon la config
	var store storage.Storer
	switch cfg.Storage.Type {
	case "memory":
		store = storage.NewMemoryStore()
		fmt.Println("📦 Utilisation du MemoryStore")
	case "json":
		store, err = storage.NewJSONStore(cfg.Storage.JSON.FilePath)
		if err != nil {
			fmt.Printf("Erreur création JSONStore: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("📄 Utilisation du JSONStore (%s)\n", cfg.Storage.JSON.FilePath)
	case "gorm":
		store, err = storage.NewGORMStore(cfg.Storage.GORM.DBPath)
		if err != nil {
			fmt.Printf("Erreur création GORMStore: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("💾 Utilisation du GORMStore (%s)\n", cfg.Storage.GORM.DBPath)
	default:
		fmt.Printf("Type de storage inconnu: %s\n", cfg.Storage.Type)
		os.Exit(1)
	}

	cmd.Execute(store)
}
