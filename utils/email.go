package utils

import (
	"regexp"
)

// checks if a string is a valid email

func IsValidEmail(email string) bool {
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regexPattern, email)
	if err != nil {
		return false
	}
	return match
}
