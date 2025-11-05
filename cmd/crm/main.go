package main

import (
	"github.com/axellelanca/2526_master2/cours2/refactor_crm_interface/internal/app"
	"github.com/axellelanca/2526_master2/cours2/refactor_crm_interface/internal/storage"
)

// La fonction main orchestre l'application. Elle dépend de l'interface Storer,
// pas de MemoryStore directement. C'est ça, l'injection de dépendances !
func main() {
	var store = storage.NewMemoryStore()
	// var store = storage.NewJSONStore()
	app.Run(store)
}
