package domain

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
	Role        Role   `json:"role"`
	Status      bool   `json:"status"`
	Avatar      string `json:"avatar"`
	CreatedAt   string `json:"created_at"`
	Users       []User `gorm:"many2many:user_cars"`
}
