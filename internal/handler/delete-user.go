package handler

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func DeleteUser(memoryStore domain.Storer, reader *bufio.Reader) (domain.Contact, error) {

	ListUsers(memoryStore)

	var contacts = memoryStore.GetAll()

	fmt.Print("Tapez l'id de l'utilisateur à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	index, err := strconv.Atoi(idStr)
	if err != nil {
		return domain.Contact{}, fmt.Errorf("index invalide")
	}

	removed := contacts[index]
	if removed == nil {
		return domain.Contact{}, fmt.Errorf("aucun contact trouvé avec l'id %d", index)
	}
	if err := memoryStore.Delete(index); err != nil {
		return domain.Contact{}, err
	}
	return *removed, nil
}
