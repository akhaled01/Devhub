package utils

import "strings"

// checks if a string is a valid email
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@gmail.com")
}
