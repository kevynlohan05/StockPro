package response

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Sector string `json:"sector"`
	Role   string `json:"role"`
	Active bool   `json:"active"`
}
