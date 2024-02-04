// hash-password
package hashpassword

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
