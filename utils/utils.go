package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass *string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*pass), bcrypt.MinCost)
	*pass = string(hash)
	return err
}