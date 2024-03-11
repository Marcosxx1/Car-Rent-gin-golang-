package userdtos

type UserInputDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Status   string `json:"status" binding:"required"`
	Avatar   string `json:"avatar"`
}

type UserUpdateDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Status string `json:"status" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

type UserOutPutDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}

type ChangePasswordDTO struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type LoginInputDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
