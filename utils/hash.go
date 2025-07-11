package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainPassword string) (string, error) {
	hashedPassBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	return string(hashedPassBytes), err
}
