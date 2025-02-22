package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	byteSlicedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(byteSlicedPassword), err
}
