package repositories

type AuthRepository interface {
	GenerateAuthToken(user_id string, user_role string) (string, error)
	ValidateAuthToken(tokenString string) error
}
