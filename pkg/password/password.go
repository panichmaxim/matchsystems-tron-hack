package password

import (
	"golang.org/x/crypto/bcrypt"
)

const Cost = 14

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, Cost)
}

func CheckPasswordHash(password, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}
