package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func DeleteUser(contacts *[]domain.Contact) (domain.Contact, error) {
	reader := bufio.NewReader(os.Stdin)

	ListUsers(contacts)

	fmt.Print("Tapez l'id de l'utilisateur à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	index, err := strconv.Atoi(idStr)
	if err != nil {
		return domain.Contact{}, fmt.Errorf("index invalide")
	}

	if index < 0 || index >= len(*contacts) {
		return domain.Contact{}, fmt.Errorf("ID indisponible")
	}

	removed := (*contacts)[index]

	*contacts = append((*contacts)[:index], (*contacts)[index+1:]...)
	return removed, nil
}
