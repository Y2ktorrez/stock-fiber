package dtos

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Telefono string `json:"telefono"`
	Rol      string `json:"rol"`
}

type UserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telefono  string `json:"telefono"`
	Rol       string `json:"rol"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
