package auth

import (
	"testing"
)

var (
	user_role = "admin"
	user_id   = "1"
)

func TestGenerateAuthTokenSuccess(t *testing.T) {
	token, err := GenerateAuthToken(user_id, user_role)
	if err != nil {
		t.Errorf("GenerateAuthToken() failed unexpectedly: %v", err)
	}

	if len(token) == 0 {
		t.Error("Generated token is empty")
	}
}

func TestGenerateAuthTokenInvalidInput(t *testing.T) {
	_, err := GenerateAuthToken("", user_role)
	if err == nil {
		t.Error("GenerateAuthToken() should fail with empty user ID")
	}

	_, err = GenerateAuthToken(user_id, "")
	if err == nil {
		t.Error("GenerateAuthToken() should fail with empty user role")
	}
}

func TestValidateAuthTokenSuccess(t *testing.T) {
	token, _ := GenerateAuthToken(user_id, user_role)

	err := ValidateAuthToken(token)
	if err != nil {
		t.Errorf("ValidateAuthToken() failed unexpectedly: %v", err)
	}
}

