package authautho

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordRepository struct{}

func NewPasswordRepository() *PasswordRepository {
	return &PasswordRepository{}
}

func (s *PasswordRepository) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *PasswordRepository) VerifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func (s *PasswordRepository) CompareHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
