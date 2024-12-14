package helpers

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) []byte {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil
	}

	return pass
}
