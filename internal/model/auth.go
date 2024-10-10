package models

type SignUpRequest struct {
	Username string  `json:"username" binding:"required"`
	Name     *string `json:"name,omitempty"`                 // Pointer to string to make it optional
	Email    string  `json:"email" binding:"required,email"` // Validates email format
}

type SignInRequest struct {
	Identifier string `json:"identifier" binding:"required"` // Can be username or email
	Password   string `json:"password" binding:"required"`   // Required password
}

type TokenResponse struct {
	Token string `json:"token" binding:"required"`
}
