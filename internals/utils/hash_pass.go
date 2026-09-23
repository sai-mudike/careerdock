package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPass(password string) (string, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashPass), fmt.Errorf("hash pass: %w", err)

}

func ComapreHashAndPass(passFromDB, passFromUser string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(passFromDB), []byte(passFromUser))

	return err == nil

}
