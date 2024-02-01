package userdtos

type UserInputDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Avatar      string `json:"avatar"`
}

type UserOutPutDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"created_at"`
}
