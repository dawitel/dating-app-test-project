package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(Password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares the hashed password from the database with the plaintext password from the request.
func CheckPassword(passwordFromReq, passwordFromDB string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(passwordFromReq))
}