package validator

import "regexp"

// IsValidZipCode verifica se o CEP possui exatamente 8 dígitos numéricos
func IsValidZipCode(zipCode string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, zipCode)
	return match
}
