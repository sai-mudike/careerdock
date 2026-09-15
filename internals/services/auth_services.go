package services

import (
	"golang.org/x/crypto/bcrypt"
)

func generateHashPass(password string) (string, error) {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashPass), err

}

func comapreHashAndPass(passFromDB, passFromUser string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(passFromDB), []byte(passFromUser))

	return err == nil

}
