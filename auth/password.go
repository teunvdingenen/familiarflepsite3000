package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(pass), err
}

func ComparePassword(password, hash string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil {
		return true
	}
	return false
}