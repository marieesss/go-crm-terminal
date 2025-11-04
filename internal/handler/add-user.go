package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func AddUser(contacts *map[int]*domain.Contact) (*domain.Contact, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(" 👀 Nom de l'utilisateur : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print(" 👀 Prénom de l'utilisateur : ")
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)

	if name == "" || surname == "" {
		return nil, fmt.Errorf("le nom et le prénom ne peuvent pas être vides")
	}

	fmt.Print(" 👀 Numéro de téléphone de l'utilisateur : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	if phone == "" {
		return nil, fmt.Errorf("le numéro de téléphone ne peut pas être vide")
	}

	fmt.Print(" 👀 Email de l'utilisateur : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if email == "" {
		return nil, fmt.Errorf("l'email ne peut pas être vide")
	}

	contact := &domain.Contact{
		Name:    name,
		Surname: surname,
		Phone:   phone,
		Email:   email,
	}

	newID := len(*contacts) + 1
	(*contacts)[newID] = contact

	return contact, nil
}
