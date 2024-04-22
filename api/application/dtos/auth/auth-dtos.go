package authdto

type ChangePasswordDTO struct {
	CurrentPassword string `json:"current_password" validate:"required,min=8"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

type LoginInputDTO struct {
	Email    string      `json:"email" validate:"required,email"`
	Password string      `json:"password" validate:"required,min=8"`
}
