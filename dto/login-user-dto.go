package dto

// LoginUserDTO is a data transfer object for login
type LoginUserDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
