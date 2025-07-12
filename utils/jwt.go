package utils

import (
	"bufio"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {
	secretKey, err := GetSigningKey()

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func GetSigningKey() (string, error) {
	file, err := os.Open("secret_val.txt")

	if err != nil {
		return "", errors.New("could not open file containing signing key")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return "", errors.New("no content found in secret file")
	}

	if scanner.Err() != nil {
		return "", errors.New("reading content from file containing signing key failed")
	}

	return scanner.Text(), nil
}
