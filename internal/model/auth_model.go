package model

type AuthRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
