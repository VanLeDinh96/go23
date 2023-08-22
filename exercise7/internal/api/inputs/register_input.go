package inputs

import "github.com/diegovanne/go23/exercise7/internal/constants/role_enums"

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role    role_enums.UserRole    `json:"role" binding:"required"`
}