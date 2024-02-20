package helpers

import "golang.org/x/crypto/bcrypt"

func ChechPasswordHash(password, hash string) bool {
	hashedPassword := []byte(hash)
	plainPassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
