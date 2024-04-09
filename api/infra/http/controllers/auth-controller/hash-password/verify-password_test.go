package hashpassword

import (
	"testing"
)

func Test_VerifyPassword(t *testing.T) {
	testPasswords := []string{
		"plainpassword",
		"AnotherPlainPassword",
		"password123!@#",
	}

	for _, plainPassword := range testPasswords {
		hashedPassword, _ := HashPassword(plainPassword)

		result := VerifyPassword(plainPassword, hashedPassword)
		if !result {
			t.Errorf("Password verification failed for password: %s", plainPassword)
		}
	}
}

func Test_VerifyPasswordFail(t *testing.T) {
	testPasswords := []string{
		"plainpassword",
		"AnotherPlainPassword",
		"password123!@#",
	}

	for _, plainPassword := range testPasswords {
		hashedPassword, _ := HashPassword(plainPassword)

		result := VerifyPassword("wrongpassword", hashedPassword)
		if result {
			t.Errorf("Password verification should have failed for password: %s", plainPassword)
		}
	}
}
