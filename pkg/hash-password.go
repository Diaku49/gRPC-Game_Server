package pkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed hashing, err: %v", err)
	}

	return string(hashPass), nil
}

func Compare(password, hashPass string) bool {
	hashbyte := []byte(hashPass)
	passbyte := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashbyte, passbyte)
	if err != nil {
		return false
	}

	return true
}
