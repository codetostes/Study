package enderecos

import (
	"slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AddressType verifies if address has a valid type
func AddressType(a string) string {
	validType := []string{"rua", "avenida", "estrada", "rodovia"}

	address1stWord := strings.Split(strings.ToLower(a), " ")[0]

	if slices.Contains(validType, address1stWord) {
		return cases.Title(language.BrazilianPortuguese).String(address1stWord)
	}
	return "Tipo inválido"
}
