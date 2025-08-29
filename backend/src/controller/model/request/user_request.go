package request

type UserRequest struct {
	Name     string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Sector   string `json:"sector" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type UserUpdateRequest struct {
	Name     string `json:"user_name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6,max=20"`
	Sector   string `json:"sector" binding:"omitempty"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user"`
	Active   bool   `json:"active" binding:"omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
