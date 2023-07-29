package utils

import "strings"

func IsValidSort(sortParam string) bool {
	validSorts := map[string]bool{
		"nome_asc":        true,
		"preco_asc":       true,
		"quantidade_asc":  true,
		"nome_desc":       true,
		"descricao_desc":  true,
		"preco_desc":      true,
		"quantidade_desc": true,
	}

	return validSorts[strings.ToLower(sortParam)]
}
