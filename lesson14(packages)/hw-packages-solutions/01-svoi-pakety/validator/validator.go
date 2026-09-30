package validator

import "strings"

// IsValidEmail - простая проверка без regexp: должна быть "@" (не первым
// и не последним символом), а после неё - домен с точкой, после которой
// есть хотя бы один символ.
func IsValidEmail(email string) bool {
	at := strings.Index(email, "@")
	if at <= 0 || at == len(email)-1 {
		return false
	}

	domain := email[at+1:]
	dot := strings.LastIndex(domain, ".")
	if dot <= 0 || dot == len(domain)-1 {
		return false
	}

	return true
}

// IsAdult проверяет совершеннолетие.
func IsAdult(age int) bool {
	return age >= 18
}
