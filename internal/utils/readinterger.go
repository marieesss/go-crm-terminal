package utils

import (
	"bufio"
	"strconv"
)

func ReadInteger(reader *bufio.Reader) int {
	id, err := ReadLine(reader)
	if err != nil {
		return -1 // Renvoie -1 pour un choix invalide
	}
	nbrId, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	return nbrId
}
