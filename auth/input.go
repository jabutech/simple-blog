package auth

// RegisterUserInput a object request for input data register
type RegisterInput struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3"`
	IsAdmin  bool   `json:"is_admin"`
}
