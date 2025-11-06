package main

import (
	"github.com/marie-mattheo/crm-project/internal/app"
	"github.com/marie-mattheo/crm-project/internal/storage"
)

// La fonction main orchestre l'application. Elle dépend de l'interface Storer,
// pas de MemoryStore directement. C'est ça, l'injection de dépendances !
func main() {
	var store = storage.NewMemoryStore()
	// var store = storage.NewJSONStore()
	app.Run(store)
}
