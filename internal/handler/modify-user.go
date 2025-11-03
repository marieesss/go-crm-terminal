package handler

import (
	"bufio"
	"fmt"
	"os"
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

func ModifyUser(contacts *[]domain.Contact) error {
	if len(*contacts) == 0 {
		return fmt.Errorf("aucun contact à modifier")
	}

	reader := bufio.NewReader(os.Stdin)

	// Afficher la liste des contacts
	fmt.Println("\n=== Liste des contacts ===")
	for i, contact := range *contacts {
		fmt.Printf("%d. %s %s\n", i+1, contact.Name, contact.Surname)
	}

	// Demander l'index du contact à modifier
	fmt.Print("\nEntrez le numéro du contact à modifier (1-" + strconv.Itoa(len(*contacts)) + ") : ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return fmt.Errorf("numéro invalide")
	}

	// Vérifier si l'index est valide
	if index < 1 || index > len(*contacts) {
		return fmt.Errorf("numéro de contact invalide")
	}

	// Ajuster l'index car les tableaux commencent à 0
	index--

	// Demander les nouvelles informations
	fmt.Print("Nouveau nom (ou appuyez sur Entrée pour garder " + (*contacts)[index].Name + ") : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		if containsNumber(name) {
			return fmt.Errorf("le nom ne doit pas contenir de chiffres")
		}
		(*contacts)[index].Name = name
	}

	fmt.Print("Nouveau prénom (ou appuyez sur Entrée pour garder " + (*contacts)[index].Surname + ") : ")
	surname, _ := reader.ReadString('\n')
	surname = strings.TrimSpace(surname)
	if surname != "" {
		if containsNumber(surname) {
			return fmt.Errorf("le prénom ne doit pas contenir de chiffres")
		}
		(*contacts)[index].Surname = surname
	}

	fmt.Print("Nouveau numéro de téléphone (ou appuyez sur Entrée pour garder " + (*contacts)[index].Phone + ") : ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	if phone != "" {
		if !validatePhoneNumber(phone) {
			return fmt.Errorf("le numéro de téléphone doit contenir uniquement des chiffres et avoir maximum 10 caractères")
		}
		(*contacts)[index].Phone = phone
	}

	fmt.Print("Nouvel email (ou appuyez sur Entrée pour garder " + (*contacts)[index].Email + ") : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		if !validateEmail(email) {
			return fmt.Errorf("l'email doit contenir un @")
		}
		(*contacts)[index].Email = email
	}

	fmt.Println("Contact modifié avec succès !")
	return nil
}
