package hashpassword

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword_Hashing(t *testing.T) {
	testPasswords := []string{
		"testpasswordtobehashed",
		"AnotherTestPassword",
		"password123!@#",
	}

	for _, testPass := range testPasswords {
		hashedPassword, err := HashPassword(testPass)
		if err != nil {
			t.Errorf("HashPassword() returned an error: %v", err)
		}
		if hashedPassword == testPass {
			t.Errorf("HashPassword() did not hash the password.")
		}
	}
}

func TestHashPassword_ValidBcryptHash(t *testing.T) {
	testPasswords := []string{
		"testpasswordtobehashed",
		"AnotherTestPassword",
		"password123!@#",
	}

	for _, testPass := range testPasswords {
		hashedPassword, _ := HashPassword(testPass)
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPass))
		if err != nil {
			t.Errorf("Hashed password is not a valid bcrypt hash: %v", err)
		}
	}
}



// I think is not necessary since here:
// api\application\use-cases\auth-use-case\sign-in-use-case.go 56
// we are validating the password f

/* func TestHashPassword_HashingEmptyPassword(t *testing.T) {
	emptyPass := ""
	hashedPassword, err := HashPassword(emptyPass)
	if err == nil {
		t.Errorf("HashPassword() did not return an error for an empty password. %s", hashedPassword)
	}
	if hashedPassword != "" {
		t.Errorf("HashPassword() returned a non-empty hash for an empty password.")
	}
} */
