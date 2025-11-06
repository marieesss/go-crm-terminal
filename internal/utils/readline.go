package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func ReadLine(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		// Affiche erreur de lecture (EOF, etc.)
		fmt.Printf("Error during read: %v\n", err)
		return "", err
	}

	trimmedInput := strings.TrimSpace(input)
	return trimmedInput, nil
}
