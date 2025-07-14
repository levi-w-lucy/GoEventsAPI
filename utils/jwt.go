package utils

import (
	"bufio"
	"errors"
	"fmt"
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

	return token.SignedString(secretKey)
}

func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// _, ok := token.Method.(*jwt.SigningMethodHMAC) //This is checking the type of Method

		// if !ok {
		// 	return nil, errors.New("Unexpected signing method")
		// }
		return GetSigningKey()
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}

func GetUserIDClaim(token *jwt.Token) (int64, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid token claims")
	}

	fmt.Println(claims["userId"])
	userId := int64(claims["userId"].(float64))

	return userId, nil
}

func GetSigningKey() ([]byte, error) {
	file, err := os.Open("secret_val.txt")

	if err != nil {
		return nil, errors.New("could not open file containing signing key")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil, errors.New("no content found in secret file")
	}

	if scanner.Err() != nil {
		return nil, errors.New("reading content from file containing signing key failed")
	}

	return scanner.Bytes(), nil
}
