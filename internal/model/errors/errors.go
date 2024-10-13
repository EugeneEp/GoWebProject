package errors

import "errors"

var (
	ErrNotFound   = errors.New("user.not_found")
	ErrNotCreated = errors.New("user.not_created")
	ErrNotDeleted = errors.New("user.not_deleted")

	ErrInvalidUserData = errors.New("auth.invalid_user_data")

	ErrInvalidAuthData = errors.New("auth.invalid_auth_data")
)
