package test

import (
	"GoWebProject/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthGetToken(t *testing.T) {
	_, err := authCase.Auth(&model.AuthRequest{
		Username: "admin",
		Password: "admin",
	})

	assert.Nil(t, err)
}
