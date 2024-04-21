package repositories

type PasswordRepository interface {
	HashPassword(password string) (string, error)
	VerifyPassword(plainPassword, hashedPassword string) bool
	CompareHashedPassword(hashedPassword, password string) error
}
