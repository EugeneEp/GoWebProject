package model

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=100"`
}

type UpdateUserRequest struct {
	UserRequest
	Password *string `json:"password" validate:"omitempty,max=100"`
	Username *string `json:"username" validate:"omitempty,max=100"`
}

type UserRequest struct {
	ID string
}
