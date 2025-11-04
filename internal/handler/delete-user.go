package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func DeleteUser(contacts *map[int]*domain.Contact) (domain.Contact, error) {
	reader := bufio.NewReader(os.Stdin)

	ListUsers(contacts)

	fmt.Print("Tapez l'id de l'utilisateur à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	index, err := strconv.Atoi(idStr)
	if err != nil {
		return domain.Contact{}, fmt.Errorf("index invalide")
	}
	valPtr, exists := (*contacts)[index]
	if !exists {
		return domain.Contact{}, fmt.Errorf("aucun contact trouvé avec cet index")
	}

	removed := *valPtr
	delete(*contacts, index)
	return removed, nil
}
