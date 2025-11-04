package handler

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/marieesss/go-crm-terminal/internal/domain"
)

func containsNumber(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func validatePhoneNumber(phone string) bool {
	if len(phone) > 10 {
		return false
	}
	for _, r := range phone {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

// Retourne le contact mis à jour + l'id choisi
func UpdateUser(memoryStore domain.Storer, reader *bufio.Reader) (*domain.Contact, int, error) {
	contacts := memoryStore.GetAll()
	if len(contacts) == 0 {
		return nil, 0, fmt.Errorf("aucun contact à modifier")
	}

	// Afficher la liste
	fmt.Println("\n=== Liste des contacts ===")
	for id, c := range contacts {
		fmt.Printf("%d) %s %s - %s - %s\n", id, c.Name, c.Surname, c.Phone, c.Email)
	}

	// Demander l'ID du contact
	fmt.Print("\nEntrez l'ID du contact à modifier : ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, 0, fmt.Errorf("id invalide")
	}
	current, ok := contacts[id]
	if !ok || current == nil {
		return nil, 0, fmt.Errorf("aucun contact trouvé avec l'id %d", id)
	}

	fmt.Printf(" 👀 Nom [%s] : ", current.Name)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = current.Name
	} else if containsNumber(name) {
		return nil, 0, fmt.Errorf("le nom ne doit pas contenir de chiffres")
	}

	fmt.Printf(" 👀 Prénom [%s] : ", current.Surname)
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)
	if surname == "" {
		surname = current.Surname
	} else if containsNumber(surname) {
		return nil, 0, fmt.Errorf("le prénom ne doit pas contenir de chiffres")
	}

	fmt.Printf(" 👀 Téléphone [%s] : ", current.Phone)
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	if phone == "" {
		phone = current.Phone
	} else if !validatePhoneNumber(phone) {
		return nil, 0, fmt.Errorf("le numéro de téléphone doit contenir uniquement des chiffres et avoir maximum 10 caractères")
	}

	fmt.Printf(" 👀 Email [%s] : ", current.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		email = current.Email
	} else if !validateEmail(email) {
		return nil, 0, fmt.Errorf("l'email doit contenir un @")
	}

	updated := &domain.Contact{
		Name:    name,
		Surname: surname,
		Phone:   phone,
		Email:   email,
	}

	if err := memoryStore.Update(id, updated); err != nil {
		return nil, 0, err
	}

	fmt.Println("✅ Contact modifié avec succès !")
	return updated, id, nil
}
