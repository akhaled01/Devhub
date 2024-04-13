package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// This function hashes a given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// This function takes in a plaintext string,
// then compares it with a hash to check if the
// hash === plaintext
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
