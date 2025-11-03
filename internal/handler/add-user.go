package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func AddUser(contacts *[]domain.Contact) (domain.Contact, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nom de l'utilisateur : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Prénom de l'utilisateur : ")
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)

	fmt.Print("Numéro de téléphone de l'utilisateur : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Email de l'utilisateur : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	contact := domain.Contact{
		Name:    name,
		Surname: surname,
		Phone:   phone,
		Email:   email,
	}
	*contacts = append(*contacts, contact)
	return contact, nil
}
