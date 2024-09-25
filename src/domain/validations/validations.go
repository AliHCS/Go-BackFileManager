package validations

import "regexp"

// ValidateEmail valida el formato del correo electrónico
func ValidateEmail(email string) bool {
	// Regex básica para validar un correo electrónico
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
