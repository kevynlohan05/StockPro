package request

type UserRequest struct {
	Name     string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Sector   string `json:"sector" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}
